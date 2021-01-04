package main

import (
	context "context"
	"fmt"
	"github.com/eibrunorodrigues/infra-grpc/orchestrator"
	"log"
	"net"
	"strings"

	"github.com/eibrunorodrigues/infra-grpc/receiver"
	"github.com/eibrunorodrigues/infra-grpc/sender"
	"github.com/eibrunorodrigues/rabbitmq-go/enums"
	"github.com/eibrunorodrigues/rabbitmq-go/rabbitmq"
	"github.com/eibrunorodrigues/rabbitmq-go/types"
	"github.com/segmentio/ksuid"
	"google.golang.org/grpc"
)

//Orchestrator struct
type Orchestrator struct {
	Client       *rabbitmq.Client
	instantiated bool
}

//Publisher struct
type Publisher struct {
	Client       *rabbitmq.Client
	instantiated bool
}

//Receiver struct
type Receiver struct {
	Client       *rabbitmq.Client
	instantiated bool
}

//New creates a new instance of a Receiver
func (r *Receiver) New() {
	client := rabbitmq.Client{}
	r.Client = &client
	r.Client.Connect()
	r.instantiated = true
}

//New creates a new instance of a Publisher
func (p *Publisher) New() {
	client := rabbitmq.Client{}
	p.Client = &client
	p.Client.Connect()
	p.instantiated = true
}

//New creates a new instance of a Orchestrator
func (o *Orchestrator) New() {
	client := rabbitmq.Client{}
	o.Client = &client
	o.Client.Connect()
	o.instantiated = true
}

//AcknowledgeMessage
func (o *Orchestrator) AcknowledgeMessage(ctx context.Context, arg *orchestrator.MessageId) (*orchestrator.Status, error) {
	if !o.instantiated {
		o.New()
	}

	err := o.Client.AcknowledgeMessage(int(arg.MessageId))

	if err != nil {
		return &orchestrator.Status{Status: false}, err
	}

	return &orchestrator.Status{Status: true}, nil
}

//RejectMessage
func (o *Orchestrator) RejectMessage(ctx context.Context, arg *orchestrator.Reject) (*orchestrator.Status, error) {
	if !o.instantiated {
		o.New()
	}

	err := o.Client.RejectMessage(int(arg.Id.MessageId), arg.Requeue)

	if err != nil {
		return &orchestrator.Status{Status: false}, err
	}

	return &orchestrator.Status{Status: true}, nil
}

//CheckIfQueueExists
func (o *Orchestrator) CheckIfQueueExists(ctx context.Context, arg *orchestrator.Name) (*orchestrator.Status, error) {
	if !o.instantiated {
		o.New()
	}

	return &orchestrator.Status{Status: o.Client.CheckIfQueueExists(arg.Name)}, nil
}

//CheckIfRouterExists
func (o *Orchestrator) CheckIfRouterExists(ctx context.Context, arg *orchestrator.Name) (*orchestrator.Status, error) {
	if !o.instantiated {
		o.New()
	}

	return &orchestrator.Status{Status: o.Client.CheckIfRouterExists(arg.Name)}, nil
}

//CreateQueue
func (o *Orchestrator) CreateQueue(ctx context.Context, arg *orchestrator.QueueCreation) (*orchestrator.Name, error) {
	if !o.instantiated {
		o.New()
	}

	queueName, err := o.Client.CreateQueue(arg.QueueName, arg.CreateDlq, arg.IsAnExclusive)
	if err != nil {
		return &orchestrator.Name{}, err
	}

	return &orchestrator.Name{Name: queueName}, nil
}

//CreateRouter
func (o *Orchestrator) CreateRouter(ctx context.Context, arg *orchestrator.RouterCreation) (*orchestrator.Name, error) {
	if !o.instantiated {
		o.New()
	}

	routerPrefix, err := enums.ParseRouterPrefix(strings.Replace(arg.Prefix.String(), "P_", "", -1))
	if err != nil {
		return &orchestrator.Name{}, err
	}

	routerType, err := enums.ParseRouterType(strings.Replace(arg.RouterType.String(), "T_", "", -1))
	if err != nil {
		return &orchestrator.Name{}, err
	}

	routerName, err := o.Client.CreateRouter(arg.RouterName, routerPrefix, routerType)
	if err != nil {
		return &orchestrator.Name{}, err
	}

	return &orchestrator.Name{Name: routerName}, nil
}

//DeleteQueue
func (o *Orchestrator) DeleteQueue(ctx context.Context, arg *orchestrator.Name) (*orchestrator.Status, error) {
	if !o.instantiated {
		o.New()
	}

	status, err := o.Client.DeleteQueue(arg.Name)
	if err != nil {
		return &orchestrator.Status{}, err
	}

	return &orchestrator.Status{Status: status}, nil
}

//DeleteRouter
func (o *Orchestrator) DeleteRouter(ctx context.Context, arg *orchestrator.Name) (*orchestrator.Status, error) {
	if !o.instantiated {
		o.New()
	}

	status, err := o.Client.DeleteRouter(arg.Name)
	if err != nil {
		return &orchestrator.Status{}, err
	}

	return &orchestrator.Status{Status: status}, nil
}

//BindQueueToRouter
func (o *Orchestrator) BindQueueToRouter(ctx context.Context, arg *orchestrator.Bind) (*orchestrator.Status, error) {
	if !o.instantiated {
		o.New()
	}

	status, err := o.Client.BindQueueToRouter(arg.Source, arg.Destination, arg.Filters)
	if err != nil {
		return &orchestrator.Status{}, err
	}

	return &orchestrator.Status{Status: status}, nil
}

//BindRouterToRouter
func (o *Orchestrator) BindRouterToRouter(ctx context.Context, arg *orchestrator.Bind) (*orchestrator.Status, error) {
	if !o.instantiated {
		o.New()
	}

	status, err := o.Client.BindRouterToRouter(arg.Source, arg.Destination, arg.Filters)
	if err != nil {
		return &orchestrator.Status{}, err
	}

	return &orchestrator.Status{Status: status}, nil
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

//Receive (*ReceiverArgs, Receiver_ReceiveServer) error
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
	var orchestrate Orchestrator
	sender.RegisterSenderServer(grpcServer, &publisher)
	receiver.RegisterReceiverServer(grpcServer, &receptor)
	orchestrator.RegisterOrchestratorServer(grpcServer, &orchestrate)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve GRPC 5052: %v", err)
	}
}
