package grpcApp

import (
	"fmt"
	grpcHandlers "github.com/weeweeshka/tataisk/internal/grpcHandlers"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

type GRPCServer struct {
	gRPCServer *grpc.Server
	port       string
	logr       *zap.Logger
}

func NewGRPCServer(port string, logr *zap.Logger, tataiskService grpcHandlers.Tataisk) *GRPCServer {
	grpcServer := grpc.NewServer()
	grpcHandlers.RegisterNewServer(grpcServer, tataiskService)
	return &GRPCServer{
		gRPCServer: grpcServer,
		port:       port,
		logr:       logr,
	}

}

func (s *GRPCServer) MustRun() error {
	l, err := net.Listen("tcp", ":"+s.port)
	if err != nil {
		return fmt.Errorf("could not listen on port %s: %w", s.port, err)
	}

	if err := s.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("could not start grpc server: %w", err)
	}

	return nil
}

func (s *GRPCServer) GracefulStop() error {
	s.gRPCServer.GracefulStop()
	return nil
}
