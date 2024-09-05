package main

import (
	"fmt"
	"grpcserver/pb"
	"io"
	"net"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedGreeterServer
}

//func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
//	return &pb.HelloResponse{Reply: "Hello" + in.Name}, nil
//}

// 服务端流式 RPC
//func (s *server) LotsOfReplies(in *pb.HelloRequest, stream pb.Greeter_LotsOfRepliesServer) error {
//	words := []string{"你好", "hello", "こんにちは", "안녕하세요"}
//	for _, word := range words {
//		data := &pb.HelloResponse{
//			Reply: word + in.GetName(),
//		}
//		// 使用Send方法返回多个数据
//		if err := stream.Send(data); err != nil {
//			return err
//		}
//	}
//	return nil
//}

// 客户端流式 RPC
//func (s *server) LotsOfGreetings(stream pb.Greeter_LotsOfGreetingsServer) error {
//	reply := "你好："
//	for {
//		// 接收客户端发来的流式数据
//		in, err := stream.Recv()
//		if err == io.EOF {
//			// 最终统一回复
//			return stream.SendAndClose(&pb.HelloResponse{
//				Reply: reply,
//			})
//		}
//		if err != nil {
//			return err
//		}
//		reply += in.GetName()
//	}
//}

// 双流式数据
//func (s *server) BindHello(stream pb.Greeter_BindHelloServer) error {
//	for {
//		// 接收流式请求
//		in, err := stream.Recv()
//		if err == io.EOF {
//			return nil
//		}
//		if err != nil {
//			return err
//		}
//		reply := magic(in.GetName())
//		if err := stream.Send(&pb.HelloResponse{Reply: reply}); err != nil {
//			return err
//		}
//	}
//}
//
//func magic(s string) string {
//	s = strings.ReplaceAll(s, "吗", "")
//	s = strings.ReplaceAll(s, "吧", "")
//	s = strings.ReplaceAll(s, "你", "我")
//	s = strings.ReplaceAll(s, "？", "!")
//	s = strings.ReplaceAll(s, "?", "!")
//	return s
//}

// SayHello 普通 metadata
//
//	func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
//		defer func() {
//			trailer := metadata.Pairs("timestamp", strconv.Itoa(int(time.Now().Unix())))
//			grpc.SetTrailer(ctx, trailer)
//		}()
//		// 从客户端请求上下文获取 metadata 数据
//		incomingContext, ok := metadata.FromIncomingContext(ctx)
//		if !ok {
//			return nil, status.Errorf(codes.DataLoss, "UnarySayHello: failed to get metadata")
//		}
//		if t, ok := incomingContext["token"]; ok {
//			fmt.Printf("token from metadata: \n")
//			if len(t) < 1 || t[0] != "token1234" {
//				return nil, status.Errorf(codes.Unimplemented, "未认证")
//			}
//		}
//
//		// 创建和发送 metadata
//		header := metadata.New(map[string]string{"location": "henan"})
//		grpc.SetHeader(ctx, header)
//
//		fmt.Printf("request received: %v, say hello...\n", in)
//
//		return &pb.HelloResponse{Reply: in.Name}, nil
//	}

func (s *server) BindHello(stream pb.Greeter_BindHelloServer) error {
	// 在defer中创建trailer记录函数的返回时间.
	defer func() {
		trailer := metadata.Pairs("timestamp", strconv.Itoa(int(time.Now().Unix())))
		stream.SetTrailer(trailer)
	}()

	// 从client读取metadata.
	md, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		return status.Errorf(codes.DataLoss, "BidirectionalStreamingSayHello: failed to get metadata")
	}

	if t, ok := md["token"]; ok {
		fmt.Printf("token from metadata:\n")
		for i, e := range t {
			fmt.Printf(" %d. %s\n", i, e)
		}
	}

	// 创建和发送header.
	header := metadata.New(map[string]string{"location": "X2Q"})
	stream.SendHeader(header)

	// 读取请求数据发送响应数据.
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		fmt.Printf("request received %v, sending reply\n", in)
		if err := stream.Send(&pb.HelloResponse{Reply: in.Name}); err != nil {
			return err
		}
	}
}
func main() {
	listen, err := net.Listen("tcp", ":9091")
	if err != nil {
		fmt.Printf("failed to listen, err:%v\n", err)
		return
	}
	g := grpc.NewServer()
	pb.RegisterGreeterServer(g, &server{})
	err = g.Serve(listen)
	if err != nil {
		fmt.Printf("failed to serve, err:%v\n", err)
		return
	}
}
