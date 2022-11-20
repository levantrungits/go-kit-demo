package transports

import (
	"context"

	kitgrpc "github.com/go-kit/kit/transport/grpc"
	kitlog "github.com/go-kit/log"
	"github.com/levantrungits/go-kit-demo/endpoints"
	"github.com/levantrungits/go-kit-demo/pb"
)

type gRPCServer struct {
	add      kitgrpc.Handler
	subtract kitgrpc.Handler
	multiply kitgrpc.Handler
	divide   kitgrpc.Handler
}

// NewGRPCServer initializes a new gRPC server
func NewGRPCServer(endpoints endpoints.MathEndpoints, logger kitlog.Logger) pb.MathServiceServer {
	return &gRPCServer{
		add: kitgrpc.NewServer(
			endpoints.Add,
			decodeMathRequest,
			encodeMathResponse,
		),
		subtract: kitgrpc.NewServer(
			endpoints.Subtract,
			decodeMathRequest,
			encodeMathResponse,
		),
		multiply: kitgrpc.NewServer(
			endpoints.Multiply,
			decodeMathRequest,
			encodeMathResponse,
		),
		divide: kitgrpc.NewServer(
			endpoints.Divide,
			decodeMathRequest,
			encodeMathResponse,
		),
	}
}

func (s *gRPCServer) Add(ctx context.Context, req *pb.MathRequest) (*pb.MathResponse, error) {
	_, resp, err := s.add.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*pb.MathResponse), nil
}

func (s *gRPCServer) Subtract(ctx context.Context, req *pb.MathRequest) (*pb.MathResponse, error) {
	_, resp, err := s.subtract.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*pb.MathResponse), nil
}

func (s *gRPCServer) Multiply(ctx context.Context, req *pb.MathRequest) (*pb.MathResponse, error) {
	_, resp, err := s.multiply.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*pb.MathResponse), nil
}

func (s *gRPCServer) Divide(ctx context.Context, req *pb.MathRequest) (*pb.MathResponse, error) {
	_, resp, err := s.divide.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*pb.MathResponse), nil
}

func decodeMathRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.MathRequest)
	return endpoints.MathReq{NumA: req.NumA, NumB: req.NumB}, nil
}

func encodeMathResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.MathResp)
	return &pb.MathResponse{Result: resp.Result}, nil
}
