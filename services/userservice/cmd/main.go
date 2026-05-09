package main

import (
	authpb "myapp/api/gen/auth"
	"myapp/pkg/config"
	"myapp/pkg/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	cfg "userapp/internal/config"
	authclient "userapp/internal/delivery/grpc/auth"
	"userapp/internal/repository/migrator"
	"userapp/internal/repository/mysql"
	"userapp/internal/repository/mysql/mysqluser"
	"userapp/internal/validator"

	"userapp/internal/delivery/grpc"
	"userapp/internal/service"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	grpc2 "google.golang.org/grpc"
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

	MysqlRepo := mysql.New(cfg2.Mysql)
	userRepo := mysqluser.New(MysqlRepo)

	conn, dErr := grpc2.Dial(cfg2.GrpcServer.AuthAddress, grpc2.WithInsecure())
	if dErr != nil {
		logger.Fatal("cannot connect to auth service: %v", zap.Error(dErr))
	}
	authClient := authpb.NewAuthServiceClient(conn)
	grpcAuthClient := authclient.NewGRPCAuthClient(authClient)

	userSvc := userservice.New(grpcAuthClient, userRepo)

	userV := validator.New(userRepo)

	grpcServer := grpc.NewServer(userV, userSvc, cfg2.GrpcServer.UserAddress)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":9091", nil)
	}()

	go func() {
		logger.Info("🚀 gRPC server starting on",
			zap.String("address", cfg2.GrpcServer.UserAddress))
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

	if err := MysqlRepo.Conn().Close(); err != nil {
		logger.Error("failed to close mysql connection", zap.Error(err))
	}

	logger.Info("Graceful shutdown completed successfully. 🛑")
}
