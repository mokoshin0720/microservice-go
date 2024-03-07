package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/mokoshin0720/microservice-go/go-kit/services/user/endpoints"
	"github.com/mokoshin0720/microservice-go/go-kit/services/user/pb"
	"github.com/mokoshin0720/microservice-go/go-kit/services/user/service"
	"github.com/mokoshin0720/microservice-go/go-kit/services/user/transport"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	fmt.Println("Hello, World!")

	var logger log.Logger
	logger = log.NewJSONLogger(os.Stdout)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	userService := service.NewService(logger)
	userEndpoint := endpoints.MakeEndpoints(userService)
	grpcServer := transport.NewGRPCServer(userEndpoint, logger)

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGALRM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	grpcListener, err := net.Listen("tcp", ":9001")
	if err != nil {
		logger.Log("during", "Listen", "err", err)
		os.Exit(1)
	}

	go func() {
		baseServer := grpc.NewServer()
		reflection.Register(baseServer)
		pb.RegisterUserServiceServer(baseServer, grpcServer)
		level.Info(logger).Log("msg", "Server started successfully 🚀")
		baseServer.Serve(grpcListener)
	}()

	level.Error(logger).Log("exit", <-errs)
}
