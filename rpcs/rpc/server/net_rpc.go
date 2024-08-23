package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

func main() {
	service := new(ServiceA)
	rpc.Register(service)
	rpc.HandleHTTP()
	listen, err := net.Listen("tcp", ":9091")
	if err != nil {
		log.Fatal(err)
	}
	http.Serve(listen, nil)
}
