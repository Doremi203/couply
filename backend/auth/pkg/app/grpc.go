package app

import (
	"auth/pkg/errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log/slog"
	"net"
)

type grpcService interface {
	Register(gRPC *grpc.Server)
	Name() string
}

type grpcServer struct {
	log *slog.Logger

	server *grpc.Server
}

func newGRPCServer(log *slog.Logger) *grpcServer {
	s := grpc.NewServer()

	reflection.Register(s)

	return &grpcServer{
		server: s,
		log:    log,
	}
}

func (s *grpcServer) Register(service grpcService) {
	s.log.Info("grpc service registered", "service", service.Name())
	service.Register(s.server)
}

func (s *grpcServer) listenAndServe() error {
	s.log.Info("starting listen on", "port", grpcConfig.Port)
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcConfig.Port))
	if err != nil {
		return errors.Wrap(err, "couldn't start grpc server listener")
	}
	s.log.Info("starting grpc server on", "port", grpcConfig.Port)

	return s.server.Serve(listener)
}

func (s *grpcServer) gracefulStop() {
	s.server.GracefulStop()
}
