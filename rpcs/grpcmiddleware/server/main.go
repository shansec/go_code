package main

import (
	"fmt"
	"io"
	"middlerwareserver/pb"
	"net"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func valid(authority []string) bool {
	if len(authority) < 1 {
		return false
	}
	token := strings.TrimPrefix(authority[0], "Bearer ")
	// 执行token认证的逻辑
	return token == "some-secret-token"
}

type server struct {
	pb.UnimplementedMiddlewareServer
}

// 方法
//func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
//	return &pb.HelloResponse{Reply: "Hello" + in.Name}, nil
//}

// BindHello 双流式数据
func (s *server) BindHello(stream pb.Middleware_BindHelloServer) error {
	for {
		// 接收流式请求
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		reply := magic(in.GetName())
		if err := stream.Send(&pb.HelloResponse{Reply: reply}); err != nil {
			return err
		}
	}
}

func magic(s string) string {
	s = strings.ReplaceAll(s, "吗", "")
	s = strings.ReplaceAll(s, "吧", "")
	s = strings.ReplaceAll(s, "你", "我")
	s = strings.ReplaceAll(s, "？", "!")
	s = strings.ReplaceAll(s, "?", "!")
	return s
}

// 中间件
//func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
//	md, ok := metadata.FromIncomingContext(ctx)
//	if !ok {
//		return nil, status.Errorf(codes.InvalidArgument, "missing metadata")
//	}
//	if !valid(md["authorization"]) {
//		return nil, status.Errorf(codes.Unauthenticated, "invalid token")
//	}
//	m, err := handler(ctx, req)
//	if err != nil {
//		fmt.Printf("RPC failed with error %v\n", err)
//	}
//	return m, err
//}

// 流式拦截器
type wrappedStream struct {
	grpc.ServerStream
}

func (w *wrappedStream) RecvMsg(m interface{}) error {
	fmt.Printf("Receive a message (Type: %T) at %v", m, time.Now().Format(time.RFC3339))
	return w.ServerStream.RecvMsg(m)
}

func (w *wrappedStream) SendMsg(m interface{}) error {
	fmt.Printf("Send a message (Type: %T) at %v", m, time.Now().Format(time.RFC3339))
	return w.ServerStream.SendMsg(m)
}

func newWrapperSteam(s grpc.ServerStream) grpc.ServerStream {
	return &wrappedStream{s}
}

func streamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	md, ok := metadata.FromIncomingContext(ss.Context())
	if !ok {
		return status.Errorf(codes.InvalidArgument, "missing metadata")
	}
	if !valid(md["authorization"]) {
		return status.Errorf(codes.Unauthenticated, "invalid token")
	}

	err := handler(srv, newWrapperSteam(ss))
	if err != nil {
		fmt.Printf("RPC failed with error: %v\n", err)
	}
	return err
}

func main() {
	listen, err := net.Listen("tcp", ":9091")
	if err != nil {
		fmt.Printf("failed to listen, err:%v\n", err)
		return
	}
	creds, _ := credentials.NewServerTLSFromFile("./server.crt", "./server.key")
	g := grpc.NewServer(grpc.Creds(creds), grpc.StreamInterceptor(streamInterceptor))
	pb.RegisterMiddlewareServer(g, &server{})
	err = g.Serve(listen)
	if err != nil {
		fmt.Printf("failed to serve, err:%v\n", err)
		return
	}
}
