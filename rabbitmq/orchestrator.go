package rabbitmq

import (
	"context"
	"github.com/eibrunorodrigues/infra-grpc/rabbitmq/proto"
	"github.com/eibrunorodrigues/rabbitmq-go/enums"
	"github.com/eibrunorodrigues/rabbitmq-go/rabbitmq"
	"github.com/eibrunorodrigues/rabbitmq-go/types"
	"strings"
)

//Orchestrator struct
type Orchestrator struct {
	Client       *rabbitmq.Client
	instantiated bool
}

//New creates a new instance of a Orchestrator
func (o *Orchestrator) New() {
	client := rabbitmq.Client{}
	o.Client = &client
	o.Client.Connect()
	o.instantiated = true
}

//AcknowledgeMessage
func (o *Orchestrator) AcknowledgeMessage(_ context.Context, arg *proto.MessageId) (*proto.Status, error) {
	if !o.instantiated {
		o.New()
	}

	err := o.Client.AcknowledgeMessage(int(arg.MessageId))

	if err != nil {
		return &proto.Status{Status: false}, err
	}

	return &proto.Status{Status: true}, nil
}

//RejectMessage
func (o *Orchestrator) RejectMessage(_ context.Context, arg *proto.Reject) (*proto.Status, error) {
	if !o.instantiated {
		o.New()
	}

	err := o.Client.RejectMessage(int(arg.Id.MessageId), arg.Requeue)

	if err != nil {
		return &proto.Status{Status: false}, err
	}

	return &proto.Status{Status: true}, nil
}

//CheckIfQueueExists
func (o *Orchestrator) CheckIfQueueExists(_ context.Context, arg *proto.Name) (*proto.Status, error) {
	if !o.instantiated {
		o.New()
	}

	boolResponse := o.Client.CheckIfQueueExists(arg.Name)
	response := proto.Status{Status: boolResponse}

	return &response, nil
}

//CheckIfRouterExists
func (o *Orchestrator) CheckIfRouterExists(_ context.Context, arg *proto.Name) (*proto.Status, error) {
	if !o.instantiated {
		o.New()
	}

	return &proto.Status{Status: o.Client.CheckIfRouterExists(arg.Name)}, nil
}

//CreateQueue
func (o *Orchestrator) CreateQueue(_ context.Context, arg *proto.QueueCreation) (*proto.Name, error) {
	if !o.instantiated {
		o.New()
	}

	queueName, err := o.Client.CreateQueue(arg.QueueName, arg.CreateDlq, arg.IsExclusive)
	if err != nil {
		return &proto.Name{}, err
	}

	return &proto.Name{Name: queueName}, nil
}

//CreateRouter
func (o *Orchestrator) CreateRouter(_ context.Context, arg *proto.RouterCreation) (*proto.Name, error) {
	if !o.instantiated {
		o.New()
	}

	routerPrefix, err := enums.ParseRouterPrefix(strings.Replace(arg.Prefix.String(), "P_", "", -1))
	if err != nil {
		return &proto.Name{}, err
	}

	routerType, err := enums.ParseRouterType(strings.Replace(arg.RouterType.String(), "T_", "", -1))
	if err != nil {
		return &proto.Name{}, err
	}

	routerName, err := o.Client.CreateRouter(arg.RouterName, routerPrefix, routerType)
	if err != nil {
		return &proto.Name{}, err
	}

	return &proto.Name{Name: routerName}, nil
}

//DeleteQueue
func (o *Orchestrator) DeleteQueue(_ context.Context, arg *proto.Name) (*proto.Status, error) {
	if !o.instantiated {
		o.New()
	}

	status, err := o.Client.DeleteQueue(arg.Name)
	if err != nil {
		return &proto.Status{}, err
	}

	return &proto.Status{Status: status}, nil
}

//DeleteRouter
func (o *Orchestrator) DeleteRouter(_ context.Context, arg *proto.Name) (*proto.Status, error) {
	if !o.instantiated {
		o.New()
	}

	status, err := o.Client.DeleteRouter(arg.Name)
	if err != nil {
		return &proto.Status{}, err
	}

	return &proto.Status{Status: status}, nil
}

//BindQueueToRouter
func (o *Orchestrator) BindQueueToRouter(_ context.Context, arg *proto.Bind) (*proto.Status, error) {
	if !o.instantiated {
		o.New()
	}

	var err error
	var status bool

	if len(arg.Filters) == 1 && arg.Filters["routing_key"] != "" {
		status, err = o.Client.BindQueueToRouter(arg.Source, arg.Destination, arg.Filters["routing_key"])
	} else {
		var filters []types.Filters
		for key, value := range arg.Filters {
			filters = append(filters, types.Filters{Key: key, Value: value})
		}

		status, err = o.Client.BindQueueToRouter(arg.Source, arg.Destination, filters)
	}

	if err != nil {
		return &proto.Status{}, err
	}

	return &proto.Status{Status: status}, nil
}

//BindRouterToRouter
func (o *Orchestrator) BindRouterToRouter(_ context.Context, arg *proto.Bind) (*proto.Status, error) {
	if !o.instantiated {
		o.New()
	}

	var err error
	var status bool

	if len(arg.Filters) == 1 && arg.Filters["routing_key"] != "" {
		status, err = o.Client.BindRouterToRouter(arg.Source, arg.Destination, arg.Filters["routing_key"])
	} else {
		var filters []types.Filters
		for key, value := range arg.Filters {
			filters = append(filters, types.Filters{Key: key, Value: value})
		}

		status, err = o.Client.BindRouterToRouter(arg.Source, arg.Destination, filters)
	}

	if err != nil {
		return &proto.Status{}, err
	}

	return &proto.Status{Status: status}, nil
}
