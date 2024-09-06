package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"middleware/pb"
	"os"
	"strings"
	"time"

	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "127.0.0.1:9091", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

// 方法
//func reply(c pb.MiddlewareClient) {
//	ctx, cancle := context.WithTimeout(context.Background(), time.Second)
//	defer cancle()
//	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
//	if err != nil {
//		log.Fatalf("could not greet: %v", err)
//	}
//	log.Printf("Greeting: %s", r.GetReply())
//}

// 双向流式数据
func runBindHello(c pb.MiddlewareClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	// 双向流模式
	stream, err := c.BindHello(ctx)
	if err != nil {
		log.Fatalf("c.BidiHello failed, err: %v", err)
	}
	waitc := make(chan struct{})
	go func() {
		for {
			// 接收服务端返回的响应
			in, err := stream.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("c.BidiHello stream.Recv() failed, err: %v", err)
			}
			fmt.Printf("AI：%s\n", in.GetReply())
		}
	}()
	// 从标准输入获取用户输入
	reader := bufio.NewReader(os.Stdin) // 从标准输入生成读对象
	for {
		cmd, _ := reader.ReadString('\n') // 读到换行
		cmd = strings.TrimSpace(cmd)
		if len(cmd) == 0 {
			continue
		}
		if strings.ToUpper(cmd) == "QUIT" {
			break
		}
		// 将获取到的数据发送至服务端
		if err := stream.Send(&pb.HelloRequest{Name: cmd}); err != nil {
			log.Fatalf("c.BidiHello stream.Send(%v) failed: %v", cmd, err)
		}
	}
	stream.CloseSend()
	<-waitc
}

// 中间件
// unaryInterceptor 客户端一元拦截器
//func unaryInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
//	var credsConfigured bool
//	for _, o := range opts {
//		_, option := o.(grpc.PerRPCCredsCallOption)
//		if option {
//			credsConfigured = true
//			break
//		}
//	}
//	if !credsConfigured {
//		opts = append(opts, grpc.PerRPCCredentials(oauth.NewOauthAccess(&oauth2.Token{
//			AccessToken: "some-secret-token",
//		})))
//	}
//	start := time.Now()
//	err := invoker(ctx, method, req, reply, cc, opts...)
//	end := time.Now()
//	fmt.Printf("RPC: %s, start time: %s, end time: %s, err: %v\n", method, start.Format("Basic"), end.Format(time.RFC3339), err)
//	return err
//}

// 流式拦截器
type wrappedStream struct {
	grpc.ClientStream
}

func (w *wrappedStream) RecvMsg(m interface{}) error {
	fmt.Printf("Receive a message (Type: %T) at %v", m, time.Now().Format(time.RFC3339))
	return w.ClientStream.RecvMsg(m)
}

func (w *wrappedStream) SendMsg(m interface{}) error {
	fmt.Printf("Send a message (Type: %T) at %v", m, time.Now().Format(time.RFC3339))
	return w.ClientStream.SendMsg(m)
}

func newWrapperSteam(s grpc.ClientStream) grpc.ClientStream {
	return &wrappedStream{s}
}

func streamInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, steamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	var credsConfigured bool
	for _, o := range opts {
		_, ok := o.(*grpc.PerRPCCredsCallOption)
		if ok {
			credsConfigured = true
			break
		}
	}
	if !credsConfigured {
		opts = append(opts, grpc.PerRPCCredentials(oauth.NewOauthAccess(&oauth2.Token{
			AccessToken: "some-secret-token",
		})))
	}
	s, err := steamer(ctx, desc, cc, method, opts...)
	if err != nil {
		return nil, err
	}
	return newWrapperSteam(s), nil
}

func main() {
	flag.Parse()

	creds, _ := credentials.NewClientTLSFromFile("./server.crt", "")
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(creds), grpc.WithStreamInterceptor(streamInterceptor))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMiddlewareClient(conn)
	runBindHello(c)
}
