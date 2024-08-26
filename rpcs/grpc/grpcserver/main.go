package main

import (
	"fmt"
	"grpcserver/pb"
	"io"
	"net"
	"strings"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreeterServer
}

//func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
//	return &pb.HelloResponse{Reply: "Hello" + in.Name}, nil
//}

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

func (s *server) BindHello(stream pb.Greeter_BindHelloServer) error {
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
