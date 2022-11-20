package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/levantrungits/go-kit-demo/service"
)

// Endpoints struct holds the list of endpoints definition
type MathEndpoints struct {
	Add      endpoint.Endpoint
	Subtract endpoint.Endpoint
	Multiply endpoint.Endpoint
	Divide   endpoint.Endpoint
}

// MathReq struct holds the endpoint request definition
type MathReq struct {
	NumA float32
	NumB float32
}

// MathResp struct holds the endpoint response definition
type MathResp struct {
	Result float32
}

// MakeEndpoints func initializes the Endpoint instance
func MakeMathEndpoints(s service.MathService) MathEndpoints {
	return MathEndpoints{
		Add:      makeAddEndpoint(s),
		Subtract: makeSubtractEndpoint(s),
		Multiply: makeMultiplyEndpoint(s),
		Divide:   makeDivideEndpoint(s),
	}
}

func makeAddEndpoint(s service.MathService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(MathReq)
		result, _ := s.Add(ctx, req.NumA, req.NumB)
		return MathResp{Result: result}, nil
	}
}

func makeSubtractEndpoint(s service.MathService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(MathReq)
		result, _ := s.Subtract(ctx, req.NumA, req.NumB)
		return MathResp{Result: result}, nil
	}
}

func makeMultiplyEndpoint(s service.MathService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(MathReq)
		result, _ := s.Multiply(ctx, req.NumA, req.NumB)
		return MathResp{Result: result}, nil
	}
}

func makeDivideEndpoint(s service.MathService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(MathReq)
		result, _ := s.Divide(ctx, req.NumA, req.NumB)
		return MathResp{Result: result}, nil
	}
}
