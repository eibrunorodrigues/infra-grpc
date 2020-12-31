package main

import (
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/eibrunorodrigues/infra-grpc/receiver"
	"github.com/eibrunorodrigues/infra-grpc/sender"
	"github.com/eibrunorodrigues/rabbitmq-go/rabbitmq"
	"github.com/eibrunorodrigues/rabbitmq-go/types"
	"github.com/segmentio/ksuid"
	"google.golang.org/grpc"
)

type Publisher struct {
	Client       *rabbitmq.Client
	instantiated bool
}

func (p *Publisher) New() {
	client := rabbitmq.Client{}
	p.Client = &client
	p.Client.Connect()
	p.instantiated = true
}

type Receiver struct {
	Client       *rabbitmq.Client
	instantiated bool
}

func (r *Receiver) New() {
	client := rabbitmq.Client{}
	r.Client = &client
	r.Client.Connect()
	r.instantiated = true
}

//Publish allows you to publish a message to RabbitMQ
func (p *Publisher) Publish(server sender.Sender_PublishServer) error {
	if !p.instantiated {
		p.New()
	}

	for {
		message, err := server.Recv()

		if err != nil {
			return err
		}

		if message.DestinyType.String() == "QUEUE" {
			if _, err := p.Client.PublishToQueue(message.Content, message.Destination, mapToFilters(message.Header)); err != nil {
				return err
			}
		} else {
			if _, err := p.Client.PublishToRouter(message.Content, message.Destination, mapToFilters(message.Header)); err != nil {
				return err
			}
		}
	}
}

//Receive(*ReceiverArgs, Receiver_ReceiveServer) error
func (r *Receiver) Receive(args *receiver.ReceiverArgs, server receiver.Receiver_ReceiveServer) error {
	if !r.instantiated {
		r.New()
	}

	fmt.Println("\nSomebody is connected")

	consumer := fmt.Sprintf("%s-%s", args.QueueName, ksuid.New().String())
	messages, err := r.Client.Connect().Consume(args.QueueName, consumer, false, false, false, false, nil)

	if err != nil {
		return err
	}

	for message := range messages {
		receiverModel := rabbitmq.GetReceiverModel(message)

		resp := receiver.ReceiverResponse{
			Filters:       filtersToMapString(receiverModel.Filters),
			Destination:   receiverModel.RouterOrigin,
			Content:       message.Body,
			IsARedelivery: receiverModel.IsARedelivery,
		}

		if err := server.Send(&resp); err != nil {
			if strings.Contains(err.Error(), "closing") {
				_ = r.Client.Close()
			} else {
				r.Client.RejectMessage(int(message.DeliveryTag), !receiverModel.IsARedelivery)
			}
			break
		}
	}
	return nil
}

func mapToFilters(items map[string]string) []types.Filters {
	var filters []types.Filters

	for key, value := range items {
		filters = append(filters, types.Filters{Key: key, Value: value})
	}

	return filters
}
func filtersToMapString(filters []types.Filters) map[string]string {
	table := make(map[string]string)
	for _, item := range filters {
		table[item.Key] = fmt.Sprintf("%v", item.Value)
	}
	return table
}

func main() {
	lis, err := net.Listen("tcp", ":5052")
	if err != nil {
		log.Fatalf("Failed to serve TCP 5052: %v", err)
	}

	grpcServer := grpc.NewServer()

	var publisher Publisher
	var receptor Receiver
	sender.RegisterSenderServer(grpcServer, &publisher)
	receiver.RegisterReceiverServer(grpcServer, &receptor)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve GRPC 5052: %v", err)
	}
}
