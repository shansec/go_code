package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"strings"

	"github.com/shansec/go_code/microservices/trimservice/pb"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 8089, "service port")
)

type Server struct {
	pb.UnimplementedTrimServer
}

func (s *Server) TrimSpace(_ context.Context, req *pb.TrimRequest) (*pb.TrimResponse, error) {
	ov := req.GetS()
	v := strings.ReplaceAll(ov, " ", "")
	fmt.Printf("ov:%s v:%v\n", ov, v)
	return &pb.TrimResponse{S: v}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}
	server := grpc.NewServer()
	pb.RegisterTrimServer(server, &Server{})
	err = server.Serve(lis)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}
