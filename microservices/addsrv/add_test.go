package main

import (
	"addsrv/pb"
	"context"
	"log"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

var bufListener *bufconn.Listener

func init() {
	bufListener = bufconn.Listen(bufSize)

	server := grpc.NewServer()
	newGRPCServer := NewGRPCServer(addService{})
	pb.RegisterAddServer(server, newGRPCServer)
	go func() {
		if err := server.Serve(bufListener); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(context.Context, string) (net.Conn, error) {
	return bufListener.Dial()
}

func TestSum(t *testing.T) {
	conn, err := grpc.DialContext(
		context.Background(),
		"bufnet",
		grpc.WithContextDialer(bufDialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewAddClient(conn)

	resp, err := c.Sum(context.Background(), &pb.SumRequest{
		A: 10,
		B: 2,
	})
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, int64(12), resp.V)
}
