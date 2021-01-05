package main

import (
	"log"
	"net"

	"github.com/eibrunorodrigues/infra-grpc/rabbitmq"
	"github.com/eibrunorodrigues/infra-grpc/rabbitmq/proto"

	"google.golang.org/grpc"
)

func initApplication() {
	lis, err := net.Listen("tcp", ":5052")
	if err != nil {
		log.Fatalf("Failed to serve TCP 5052: %v", err)
	}

	grpcServer := grpc.NewServer()

	proto.RegisterSenderServer(grpcServer, &rabbitmq.Publisher{})
	proto.RegisterReceiverServer(grpcServer, &rabbitmq.Receiver{})
	proto.RegisterOrchestratorServer(grpcServer, &rabbitmq.Orchestrator{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve GRPC 5052: %v", err)
	}
}

func main() {
	initApplication()
}
