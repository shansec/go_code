package main

import (
	"context"
	"grpcgateway/proto/gateway"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	gateway.UnimplementedGatewayServer
}

func (s *server) SayHello(ctx context.Context, in *gateway.HelloRequest) (*gateway.HelloResponse, error) {
	return &gateway.HelloResponse{Reply: in.Name + " ssss"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	gateway.RegisterGatewayServer(s, &server{})
	log.Println("Serving gRPC on 0.0.0.0:8080")
	log.Fatal(s.Serve(lis))
}
