package grpcApp

import (
	"go.uber.org/zap"
	"google.golang.org/grpc"
	grpcHandlers "github.com/weeweeshka/tataisk/internal/grpcHandlers"
)

type GRPCServer struct {
	gRPCServer *grpc.Server
	port       string
	logr       *zap.Logger
}

func NewGRPCServer(port string, logr *zap.Logger, tataiskService) *GRPCServer {}
