package main

import (
	"addsrv/pb"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/go-kit/log"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

var (
	httpAddr = flag.String("http-addr", ":8080", "http listen address")
	grpcAddr = flag.String("grpc-addr", ":8081", "gRpc listen address")
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)
	bs := NewLogMiddlewareService()
	bsMiddleware := NewLogMiddleware(logger, bs)

	var g errgroup.Group

	g.Go(func() error {
		httpListener, err := net.Listen("tcp", *httpAddr)
		if err != nil {
			fmt.Printf("http: net.Listen(tcp, %s) failed, err:%v\n", *httpAddr, err)
			return err
		}
		defer httpListener.Close()
		httpHandler := NewHTTPServer(bsMiddleware, logger)
		return http.Serve(httpListener, httpHandler)
	})

	g.Go(func() error {
		// gRPC服务
		grpcListener, err := net.Listen("tcp", *grpcAddr)
		if err != nil {
			fmt.Printf("grpc: net.Listen(tcp, %s) faield, err:%v\n", *grpcAddr, err)
			return err
		}
		defer grpcListener.Close()
		s := grpc.NewServer()
		pb.RegisterAddServer(s, NewGRPCServer(bsMiddleware))
		return s.Serve(grpcListener)
	})

	if err := g.Wait(); err != nil {
		fmt.Printf("server exit with err:%v\n", err)
	}
}
