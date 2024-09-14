package main

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/shansec/go_code/microservices/trimservice/pb"
	"google.golang.org/grpc"
)

type SumRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}

type SumResponse struct {
	V   int    `json:"v"`
	Err string `json:"err,omitempty"`
}

type ConcatRequest struct {
	A string `json:"a"`
	B string `json:"b"`
}

type ConcatResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"`
}

func makeSumEndpoint(svc AddService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SumRequest)
		sum, err := svc.Sum(ctx, req.A, req.B)
		if err != nil {
			return SumResponse{V: sum, Err: err.Error()}, nil
		}
		return SumResponse{V: sum}, nil
	}
}

func makeConcatEndpoint(svc AddService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ConcatRequest)
		con, err := svc.Concat(ctx, req.A, req.B)
		if err != nil {
			return ConcatResponse{V: con, Err: err.Error()}, nil
		}
		return ConcatResponse{V: con}, nil
	}
}

// TrimRequest 调用其它服务
type TrimRequest struct {
	s string
}

type TrimResponse struct {
	s string
}

func makeTrimEndpoint(conn *grpc.ClientConn) endpoint.Endpoint {
	return grpctransport.NewClient(
		conn,
		"pb.Trim",
		"TrimSpace",
		encodeTrimRequest,
		decodeTrimResponse,
		pb.TrimResponse{},
	).Endpoint()
}
