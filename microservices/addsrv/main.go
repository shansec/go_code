package main

import (
	"flag"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	httpAddr = flag.String("http-addr", ":8080", "http listen address")
	grpcAddr = flag.String("grpc-addr", ":8081", "gRpc listen address")
	trimAddr = flag.String("trim-addr", ":8089", "trimSpace listen address")
)

func main() {
	//logger := log.NewLogfmtLogger(os.Stderr)
	//bs := NewLogMiddlewareService()
	//bsMiddleware := NewLogMiddleware(logger, bs)
	//
	//var g errgroup.Group
	//
	//g.Go(func() error {
	//	httpListener, err := net.Listen("tcp", *httpAddr)
	//	if err != nil {
	//		fmt.Printf("http: net.Listen(tcp, %s) failed, err:%v\n", *httpAddr, err)
	//		return err
	//	}
	//	defer httpListener.Close()
	//	httpHandler := NewHTTPServer(bsMiddleware, logger)
	//	return http.Serve(httpListener, httpHandler)
	//})
	//
	//g.Go(func() error {
	//	// gRPC服务
	//	grpcListener, err := net.Listen("tcp", *grpcAddr)
	//	if err != nil {
	//		fmt.Printf("grpc: net.Listen(tcp, %s) faield, err:%v\n", *grpcAddr, err)
	//		return err
	//	}
	//	defer grpcListener.Close()
	//	s := grpc.NewServer()
	//	pb.RegisterAddServer(s, NewGRPCServer(bsMiddleware))
	//	return s.Serve(grpcListener)
	//})
	//
	//if err := g.Wait(); err != nil {
	//	fmt.Printf("server exit with err:%v\n", err)
	//}

	// 调用其它服务
	bs := NewService()
	conn, err := grpc.Dial(*trimAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return
	}
	defer conn.Close()
	trimEndpoint := makeTrimEndpoint(conn)
	bs = NewServiceWithTrim(trimEndpoint, bs)
}
