package main

import (
	"context"
	"gatewayapp/internal/client/authclient"
	"gatewayapp/internal/client/productclient"
	"gatewayapp/internal/client/storeclient"
	"gatewayapp/internal/client/userclient"
	cfg "gatewayapp/internal/config"
	httpserver "gatewayapp/internal/delivery/http"
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

	authClient, aErr := authclient.New(cfg2.GrpcClient.AuthAddress)
	if aErr != nil {
		logger.Fatal("failed to initialize auth client", zap.Error(aErr))
	}

	defer authClient.Close()

	userClient, uErr := userclient.New(cfg2.GrpcClient.UserAddress)
	if uErr != nil {
		logger.Fatal("failed to initialize user client", zap.Error(uErr))
	}

	defer userClient.Close()

	storeClient, sErr := storeclient.New(cfg2.GrpcClient.StoreAddress)
	if sErr != nil {
		logger.Fatal("failed to initialize store client", zap.Error(sErr))
	}

	defer storeClient.Close()

	productClient, pErr := productclient.New(cfg2.GrpcClient.ProductAddress)
	if pErr != nil {
		logger.Fatal("failed to initialize product client", zap.Error(pErr))
	}

	defer productClient.Close()

	server := httpserver.New(*userClient, *authClient, *storeClient, *productClient, cfg2)

	go func() {
		logger.Info("🚀 Starting Gateway Service...", zap.String("address", cfg2.HttpServer.Address))
		server.Serve()
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	logger.Info("Received shutdown signal. Initiating graceful shutdown...")

	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), cfg2.Application.GracefulShutdownTimeout)
	defer cancel()

	if err := server.Router.Shutdown(ctxWithTimeout); err != nil {
		logger.Error("graceful shutdown error", zap.Error(err))
	}

	<-ctxWithTimeout.Done()
	logger.Info("Graceful shutdown completed successfully. 🛑")
}
