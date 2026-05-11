package main

import (
	authpb "myapp/api/gen/auth"
	"myapp/pkg/config"
	"myapp/pkg/logger"
	"myapp/pkg/metrics" // ← اضافه شد
	"os"
	"os/signal"
	"syscall"
	cfg "userapp/internal/config"
	"userapp/internal/delivery/grpc"
	authclient "userapp/internal/delivery/grpc/auth"
	"userapp/internal/repository/migrator"
	"userapp/internal/repository/mysql"
	"userapp/internal/repository/mysql/mysqluser"
	userservice "userapp/internal/service"
	"userapp/internal/validator"

	"go.uber.org/zap"
	grpc2 "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure" // ← جایگزین WithInsecure
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

	mgr := migrator.New(cfg2.Mysql)
	mgr.Up()

	mysqlRepo := mysql.New(cfg2.Mysql)
	userRepo := mysqluser.New(mysqlRepo)

	conn, dErr := grpc2.NewClient(
		cfg2.GrpcServer.AuthAddress,
		grpc2.WithTransportCredentials(insecure.NewCredentials()),
	)
	if dErr != nil {
		logger.Fatal("cannot connect to auth service", zap.Error(dErr))
	}

	authClient := authpb.NewAuthServiceClient(conn)
	grpcAuthClient := authclient.NewGRPCAuthClient(authClient)

	userSvc := userservice.New(grpcAuthClient, userRepo)
	userV := validator.New(userRepo)

	grpcServer := grpc.NewServer(userV, userSvc, cfg2.GrpcServer.UserAddress)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go metrics.StartServer(cfg2.Metrics.Port)

	go func() {
		logger.Info("🚀 gRPC server starting", zap.String("address", cfg2.GrpcServer.UserAddress))
		if err := grpcServer.Run(); err != nil {
			logger.Fatal("cannot start grpc server", zap.Error(err))
		}
	}()

	<-quit
	logger.Info("Received shutdown signal. Initiating graceful shutdown...")

	grpcServer.GracefulStop()

	if err := conn.Close(); err != nil {
		logger.Error("failed to close auth service connection", zap.Error(err))
	}
	if err := mysqlRepo.Conn().Close(); err != nil {
		logger.Error("failed to close mysql connection", zap.Error(err))
	}

	logger.Info("Graceful shutdown completed. 🛑")
}
