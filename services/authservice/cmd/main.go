package main

import (
	cfg "authapp/internal/config"
	"authapp/internal/delivery/grpc"
	authservice "authapp/internal/service"
	"myapp/pkg/config"
	"myapp/pkg/logger"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
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

	authSvc := authservice.New(cfg2.Auth)

	grpcServer := grpc.NewServer(authSvc, cfg2.GrpcServer.AuthAddress, cfg2.Metrics.Port)

	go func() {
		logger.Info("🚀gRPC server started on ",
			zap.String("address", cfg2.GrpcServer.AuthAddress))

		if err := grpcServer.Run(); err != nil {
			logger.Fatal("cannot start grpc server", zap.Error(err))
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	logger.Info("Received shutdown signal. Initiating graceful shutdown...")

	grpcServer.GracefulStop()

	logger.Info("Graceful shutdown completed successfully. 🛑")

}
