package rabbitmq

import (
	"github.com/eibrunorodrigues/infra-grpc/rabbitmq/proto"
	"github.com/eibrunorodrigues/rabbitmq-go/rabbitmq"
	"github.com/eibrunorodrigues/rabbitmq-go/types"
)

//Publisher struct
type Publisher struct {
	Client       *rabbitmq.Client
	instantiated bool
}

//New creates a new instance of a Publisher
func (p *Publisher) New() {
	client := rabbitmq.Client{}
	p.Client = &client
	p.Client.Connect()
	p.instantiated = true
}


//Publish allows you to publish a message to RabbitMQ
func (p *Publisher) Publish(server proto.Sender_PublishServer) error {
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

func mapToFilters(items map[string]string) []types.Filters {
	var filters []types.Filters

	for key, value := range items {
		filters = append(filters, types.Filters{Key: key, Value: value})
	}

	return filters
}
