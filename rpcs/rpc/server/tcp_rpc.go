package main

import (
	"log"
	"net"
	"net/rpc"
)

type Args struct {
	X, Y int
}

type ServiceA struct{}

func (s *ServiceA) Add(args *Args, reply *int) error {
	*reply = args.X + args.Y
	return nil
}

func main() {
	service := new(ServiceA)
	rpc.Register(service)
	listen, err := net.Listen("tcp", ":9091")
	if err != nil {
		log.Fatal(err)
	}
	for {
		accept, _ := listen.Accept()
		rpc.ServeConn(accept)
	}
}
