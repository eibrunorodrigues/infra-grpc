module github.com/eibrunorodrigues/infra-grpc

go 1.15

require (
	github.com/eibrunorodrigues/rabbitmq-go v0.0.0-20201230225947-1feb4b5bb83e
	github.com/golang/protobuf v1.4.2
	github.com/segmentio/ksuid v1.0.3
	google.golang.org/grpc v1.34.0
)

replace github.com/eibrunorodrigues/rabbitmq-go => ../github.com/eibrunorodrigues/rabbitmq-go
