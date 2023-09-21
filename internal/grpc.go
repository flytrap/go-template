package app

import (
	"fmt"
	"net"
	"os"

	"github.com/flytrap/go-template/internal/config"
	"github.com/flytrap/go-template/pb/v1"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
)

func InitGrpcServer(logService pb.LogServiceServer) *GrpcServer {
	return &GrpcServer{config: config.C.ServerConfig, logService: logService}
}

type GrpcServer struct {
	config     config.ServerConfig
	logService pb.LogServiceServer
}

func (s *GrpcServer) Run() error {
	addr := fmt.Sprintf("%s:%s", s.config.Host, s.config.GrpcPort)
	logrus.Info(addr)
	listener, err := net.Listen(s.config.GrpcProtocol, addr)

	if err != nil {
		logrus.Error(err)
		return err
	}

	opts := []grpc.ServerOption{}

	if len(s.config.Cert) > 0 && len(s.config.Key) > 0 {
		c, err := credentials.NewServerTLSFromFile(s.config.Cert, s.config.Key)
		if err != nil {
			logrus.Warn(err)
		} else {
			opts = append(opts, grpc.Creds(c))
		}
	}

	grpcLog := grpclog.NewLoggerV2(os.Stdout, os.Stderr, os.Stderr)
	grpclog.SetLoggerV2(grpcLog)

	srv := grpc.NewServer(opts...)

	pb.RegisterLogServiceServer(srv, s.logService)

	reflection.Register(srv)

	if err := srv.Serve(listener); err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}
