package main

import (
	"context"
	"flag"
	"fmt"
	"grpcclient/pb"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "127.0.0.1:9091", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

//func reply() {
//	ctx, cancle := context.WithTimeout(context.Background(), time.Second)
//	defer cancle()
//	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
//	if err != nil {
//		log.Fatalf("could not greet: %v", err)
//	}
//	log.Printf("Greeting: %s", r.GetReply())
//}

// 服务端流式
//func runLotsOfReplies(c pb.GreeterClient) {
//	// server端流式RPC
//	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
//	defer cancel()
//	stream, err := c.LotsOfReplies(ctx, &pb.HelloRequest{Name: *name})
//	if err != nil {
//		log.Fatalf("c.LotsOfReplies failed, err: %v", err)
//	}
//	for {
//		// 接收服务端返回的流式数据，当收到io.EOF或错误时退出
//		res, err := stream.Recv()
//		if err == io.EOF {
//			break
//		}
//		if err != nil {
//			log.Fatalf("c.LotsOfReplies failed, err: %v", err)
//		}
//		log.Printf("got reply: %q\n", res.GetReply())
//	}
//}

// 客户端流式
//func LotsOfGreetings(c pb.GreeterClient) {
//	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
//	defer cancel()
//	// 客户端流式RPC
//	stream, err := c.LotsOfGreetings(ctx)
//	if err != nil {
//		log.Fatalf("c.LotsOfGreetings failed, err: %v", err)
//	}
//	names := []string{"七米", "沙河娜扎"}
//	for _, name := range names {
//		// 发送流式数据
//		err := stream.Send(&pb.HelloRequest{Name: name})
//		if err != nil {
//			log.Fatalf("c.LotsOfGreetings stream.Send(%v) failed, err: %v", name, err)
//		}
//	}
//	res, err := stream.CloseAndRecv()
//	if err != nil {
//		log.Fatalf("c.LotsOfGreetings failed: %v", err)
//	}
//	log.Printf("got reply: %v", res.GetReply())
//}

// 双向流式数据
//func runBindHello(c pb.GreeterClient) {
//	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
//	defer cancel()
//	// 双向流模式
//	stream, err := c.BindHello(ctx)
//	if err != nil {
//		log.Fatalf("c.BidiHello failed, err: %v", err)
//	}
//	waitc := make(chan struct{})
//	go func() {
//		for {
//			// 接收服务端返回的响应
//			in, err := stream.Recv()
//			if err == io.EOF {
//				// read done.
//				close(waitc)
//				return
//			}
//			if err != nil {
//				log.Fatalf("c.BidiHello stream.Recv() failed, err: %v", err)
//			}
//			fmt.Printf("AI：%s\n", in.GetReply())
//		}
//	}()
//	// 从标准输入获取用户输入
//	reader := bufio.NewReader(os.Stdin) // 从标准输入生成读对象
//	for {
//		cmd, _ := reader.ReadString('\n') // 读到换行
//		cmd = strings.TrimSpace(cmd)
//		if len(cmd) == 0 {
//			continue
//		}
//		if strings.ToUpper(cmd) == "QUIT" {
//			break
//		}
//		// 将获取到的数据发送至服务端
//		if err := stream.Send(&pb.HelloRequest{Name: cmd}); err != nil {
//			log.Fatalf("c.BidiHello stream.Send(%v) failed: %v", cmd, err)
//		}
//	}
//	stream.CloseSend()
//	<-waitc
//}

// 普通 RPC 调用 metadata
//func unaryCallWithMetadata(c pb.GreeterClient, name string) {
//	fmt.Println("--- unarySayHello Client ---")
//	// 创建 metadata
//	md := metadata.Pairs(
//		"token", "token123",
//		"request_id", "123456",
//	)
//	// 基于 metadata 创建 context
//	ctx := metadata.NewOutgoingContext(context.Background(), md)
//	var header, trailer metadata.MD
//	hello, err := c.SayHello(
//		ctx,
//		&pb.HelloRequest{Name: name},
//		grpc.Header(&header),   // 接受服务端发来的 header
//		grpc.Trailer(&trailer), // 接受服务端发来的 trailer
//	)
//	if err != nil {
//		fmt.Printf("failed to call service: %v\n", err)
//		return
//	}
//	// 从 header 中取 location
//	if t, ok := header["location"]; ok {
//		fmt.Printf("location from header: \n")
//		for i, v := range t {
//			fmt.Printf(" %d. %s\n", i, v)
//		}
//	} else {
//		fmt.Println("location get failed or location in header")
//		return
//	}
//	// 获取响应结果
//	fmt.Printf("server response: %v\n", hello.Reply)
//	// 从 trailer 中取 timestamp
//	if t, ok := trailer["timestamp"]; ok {
//		fmt.Printf("timestamp from trailer:\n")
//		for i, e := range t {
//			fmt.Printf(" %d. %s\n", i, e)
//		}
//	} else {
//		log.Printf("timestamp expected but doesn't exist in trailer")
//	}
//}

// bidirectionalWithMetadata 双向流式调用 metadata
func bidirectionalWithMetadata(c pb.GreeterClient, name string) {
	// 创建 metadata 和 context
	md := metadata.Pairs("token", "token1234")
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	stream, err := c.BindHello(ctx)
	if err != nil {
		log.Fatalf("failed to call binding hello stream: %v\n", err)
	}
	go func() {
		header, err := stream.Header()
		if err != nil {
			log.Fatalf("failed to get header from stream: %v\n", err)
		}
		// 从响应的 header 中读取数据
		if t, ok := header["location"]; ok {
			fmt.Printf("location from header: \n")
			for i, e := range t {
				fmt.Printf(" %d. %s\n", i, e)
			}
		} else {
			log.Println("location expected but doesn't exist in header")
			return
		}

		// 发送所有的请求数据到 server
		for i := 0; i < 5; i++ {
			if err := stream.Send(&pb.HelloRequest{Name: name}); err != nil {
				log.Fatalf("failed to send server: %v\n", err)
			}
		}
		stream.CloseSend()
	}()

	// 读取所有的响应
	var rpcStatus error
	fmt.Printf("got response: \n")
	for {
		recv, err := stream.Recv()
		if err != nil {
			rpcStatus = err
			break
		}
		fmt.Printf(" - %s\n", recv.Reply)
	}
	if rpcStatus != io.EOF {
		log.Printf("failed to finish server streaming: %v", rpcStatus)
		return
	}
	// 当 rpc 结束时读取 trailer
	trailer := stream.Trailer()
	if t, ok := trailer["timestamp"]; ok {
		fmt.Printf("timestamp from trailer:\n")
		for i, e := range t {
			fmt.Printf(" %d. %s\n", i, e)
		}
	} else {
		log.Printf("timestamp expected but doesn't exist in trailer")
	}
}

func main() {
	flag.Parse()
	// 连接 server 端，此处禁止用安全输入
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	bidirectionalWithMetadata(c, "五月")
}
