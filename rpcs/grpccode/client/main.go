package main

import (
	"context"
	"flag"
	"fmt"
	"grpccodeclient/pb"
	"log"
	"time"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

var name = flag.String("name", "五月", "通过-name告诉server你是谁")

func main() {
	flag.Parse() // 解析命令行参数

	// 连接server
	conn, err := grpc.Dial("127.0.0.1:8972", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("grpc.Dial failed,err:%v", err)
		return
	}
	defer conn.Close()
	// 创建客户端
	c := pb.NewGreeterClient(conn) // 使用生成的Go代码
	// 调用RPC方法
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		s := status.Convert(err)        // 将err转为status
		for _, d := range s.Details() { // 获取details
			switch info := d.(type) {
			case *errdetails.QuotaFailure:
				fmt.Printf("Quota failure: %s\n", info)
			default:
				fmt.Printf("Unexpected type: %s\n", info)
			}
		}
		fmt.Printf("c.SayHello failed, err:%v\n", err)
		return
	}
	// 拿到了RPC响应
	log.Printf("resp:%v\n", resp.GetReply())
}
