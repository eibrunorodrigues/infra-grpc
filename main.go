package main

import (
	"github.com/eibrunorodrigues/infra-grpc/rabbitmq"
	"github.com/eibrunorodrigues/infra-grpc/rabbitmq/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)


func main() {
	lis, err := net.Listen("tcp", ":5052")
	if err != nil {
		log.Fatalf("Failed to serve TCP 5052: %v", err)
	}

	grpcServer := grpc.NewServer()

	var publisher rabbitmq.Publisher
	var receptor rabbitmq.Receiver
	var orchestrate rabbitmq.Orchestrator
	proto.RegisterSenderServer(grpcServer, &publisher)
	proto.RegisterReceiverServer(grpcServer, &receptor)
	proto.RegisterOrchestratorServer(grpcServer, &orchestrate)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve GRPC 5052: %v", err)
	}
}
