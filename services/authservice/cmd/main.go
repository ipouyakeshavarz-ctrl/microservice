package main

import (
	cfg "authapp/internal/config"
	authservice "authapp/internal/service"
	gen "myapp/api/gen/auth"
	"myapp/pkg/config"
	"myapp/pkg/logger"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	var cfg2 cfg.Config
	err := config.Load("config.yml", &cfg2)
	if err != nil {
		panic("failed to load config: " + err.Error())
	}

	logger.InitLogger(cfg2.Logger.ServiceName, cfg2.Logger.Development, cfg2.Logger.FilePath)
	defer logger.Sync()

	logger.Info("config", zap.Any("config", cfg2))
	lis, lErr := net.Listen("tcp", cfg2.GrpcServer.AuthAddress)
	if lErr != nil {
		logger.Fatal("failed to listen:", zap.Error(lErr))
	}

	grpcServer := grpc.NewServer()

	authServer := authservice.New(cfg2.Auth)

	gen.RegisterAuthServiceServer(grpcServer, authServer)

	logger.Info("🚀 AuthService gRPC server running on ",
		zap.String("address:", cfg2.GrpcServer.AuthAddress))

	if err := grpcServer.Serve(lis); err != nil {
		logger.Fatal("failed to serve:", zap.Error(err))
	}
}
