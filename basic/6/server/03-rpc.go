package main

import (
	"net"
	"net/http"
	"net/rpc"
)

type Server struct {
}

type Req struct {
	NumOne int
	NumTwo int
}

type Res struct {
	Sum int
}

func (s *Server) Add(req Req, res *Res) error {
	res.Sum = req.NumOne + req.NumTwo
	return nil
}

func main() {
	rpc.Register(new(Server))
	rpc.HandleHTTP()
	l, _ := net.Listen("tcp", ":8888")
	http.Serve(l, nil)
}
