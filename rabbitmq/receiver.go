package rabbitmq

import (
	"fmt"
	"github.com/eibrunorodrigues/infra-grpc/rabbitmq/proto"
	"github.com/eibrunorodrigues/rabbitmq-go/rabbitmq"
	"github.com/eibrunorodrigues/rabbitmq-go/types"
	"github.com/segmentio/ksuid"
	"strings"
)

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

//Receive (*ReceiverArgs, Receiver_ReceiveServer) error
func (r *Receiver) Receive(args *proto.ReceiverArgs, server proto.Receiver_ReceiveServer) error {
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

		resp := proto.ReceiverResponse{
			Filters:       filtersToMapString(receiverModel.Filters),
			Destination:   receiverModel.RouterOrigin,
			Content:       message.Body,
			IsARedelivery: receiverModel.IsARedelivery,
			MessageId:     int64(receiverModel.MessageId),
		}

		if err := server.Send(&resp); err != nil {
			if strings.Contains(err.Error(), "closing") {
				_ = r.Client.Close()
			} else {
				_ = r.Client.RejectMessage(int(message.DeliveryTag), !receiverModel.IsARedelivery)
			}
			break
		}
	}
	return nil
}

func filtersToMapString(filters []types.Filters) map[string]string {
	table := make(map[string]string)
	for _, item := range filters {
		table[item.Key] = fmt.Sprintf("%v", item.Value)
	}
	return table
}
