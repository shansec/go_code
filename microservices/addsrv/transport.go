package main

import (
	"addsrv/pb"
	"context"
	"encoding/json"
	"net/http"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func decodeGRPCSumRequest(_ context.Context, gRpcReq interface{}) (interface{}, error) {
	req := gRpcReq.(*pb.SumRequest)
	return SumRequest{A: int(req.A), B: int(req.B)}, nil
}

func decodeGRPCConcatRequest(_ context.Context, gRpcReq interface{}) (interface{}, error) {
	req := gRpcReq.(*pb.ConcatRequest)
	return ConcatRequest{A: req.A, B: req.B}, nil
}

func encodeGRPCSumResponse(_ context.Context, gRpcRes interface{}) (interface{}, error) {
	res := gRpcRes.(SumResponse)
	return &pb.SumResponse{V: int64(res.V), Err: res.Err}, nil
}

func encodeGRPCConcatResponse(_ context.Context, gRpcRes interface{}) (interface{}, error) {
	res := gRpcRes.(ConcatResponse)
	return &pb.ConcatResponse{V: res.V, Err: res.Err}, nil
}

type grpcServer struct {
	pb.UnimplementedAddServer
	sum    grpctransport.Handler
	concat grpctransport.Handler
}

func (s *grpcServer) Sum(ctx context.Context, req *pb.SumRequest) (*pb.SumResponse, error) {
	_, res, err := s.sum.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.SumResponse), nil
}

func (s *grpcServer) Concat(ctx context.Context, req *pb.ConcatRequest) (*pb.ConcatResponse, error) {
	_, res, err := s.concat.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.ConcatResponse), nil
}

func NewGRPCServer(svc AddService) pb.AddServer {
	return &grpcServer{
		sum:    grpctransport.NewServer(makeSumEndpoint(svc), decodeGRPCSumRequest, encodeGRPCSumResponse),
		concat: grpctransport.NewServer(makeConcatEndpoint(svc), decodeGRPCConcatRequest, encodeGRPCConcatResponse),
	}
}

func decodeSumRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req SumRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}
func decodeConcatRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req ConcatRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func NewHTTPServer(svc AddService) http.Handler {
	sumHandler := httpTransport.NewServer(makeSumEndpoint(svc), decodeSumRequest, encodeResponse)
	concatHandler := httpTransport.NewServer(makeConcatEndpoint(svc), decodeConcatRequest, encodeResponse)

	// use github.com/gorilla/mux
	r := mux.NewRouter()
	r.Handle("/sum", sumHandler).Methods("POST")
	r.Handle("/concat", concatHandler).Methods("POST")

	// use gin
	// r := gin.Default()
	// r.POST("/sum", gin.WrapH(sumHandler))
	// r.POST("/concat", gin.WrapH(concatHandler))
	return r
}
