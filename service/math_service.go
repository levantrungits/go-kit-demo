package service

import (
	"context"

	kitlog "github.com/go-kit/log"
)

type mathservice struct {
	logger kitlog.Logger
}

// Service Interface
type MathService interface {
	Add(ctx context.Context, numA, numB float32) (float32, error)
	Subtract(ctx context.Context, numA, numB float32) (float32, error)
	Multiply(ctx context.Context, numA, numB float32) (float32, error)
	Divide(ctx context.Context, numA, numB float32) (float32, error)
}

// NewMathService func initializes a service
func NewMathService(logger kitlog.Logger) MathService {
	return &mathservice{
		logger: logger,
	}
}

func (s mathservice) Add(ctx context.Context, numA, numB float32) (float32, error) {
	return numA + numB, nil
}

func (s mathservice) Subtract(ctx context.Context, numA, numB float32) (float32, error) {
	return numA - numB, nil
}

func (s mathservice) Multiply(ctx context.Context, numA, numB float32) (float32, error) {
	return numA * numB, nil
}

func (s mathservice) Divide(ctx context.Context, numA, numB float32) (float32, error) {
	return numA / numB, nil
}
