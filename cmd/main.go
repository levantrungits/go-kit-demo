package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/levantrungits/go-kit-demo/endpoints"
	"github.com/levantrungits/go-kit-demo/pb"
	"github.com/levantrungits/go-kit-demo/service"
	"github.com/levantrungits/go-kit-demo/transports"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// initializes a Logger
	var logger log.Logger
	logger = log.NewJSONLogger(os.Stdout)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	// Initializes a service
	addservice := service.NewMathService(logger)
	// Initializes the Endpoint instance
	addendpoint := endpoints.MakeMathEndpoints(addservice)
	// Initializes a new gRPC server
	grpcServer := transports.NewGRPCServer(addendpoint, logger)

	// Set up channel on which to send signal notifications.
	// We must use a buffered channel or risk missing the signal
	// if we're not ready to receive when the signal is sent.
	errs := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		// Passing no signals to Notify means that
		// all signals will be sent to the channel.
		signal.Notify(c)
		// Block until any signal is received.
		errs <- fmt.Errorf("%s", <-c)
	}()

	// Initializes a listener gRPC
	grpcListener, err := net.Listen("tcp", ":50051")
	if err != nil {
		logger.Log("during", "Listen", "err", err)
		os.Exit(1)
	}

	// New gRPC server
	go func() {
		baseServer := grpc.NewServer()
		// Register gRPC service
		pb.RegisterMathServiceServer(baseServer, grpcServer)
		level.Info(logger).Log("msg", "gRPC Server started successfully 🚀")
		// Package reflection implements server reflection service.
		reflection.Register(baseServer)
		// Launch servicing gRPC
		baseServer.Serve(grpcListener)
	}()

	level.Error(logger).Log("exit", <-errs)
}
