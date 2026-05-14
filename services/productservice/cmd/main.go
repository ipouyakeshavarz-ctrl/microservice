package main

import (
	"context"
	"myapp/pkg/config"
	"myapp/pkg/logger"
	"os"
	"os/signal"
	cfg "productapp/internal/config"
	"productapp/internal/delivery/grpc"
	"productapp/internal/repository/migrator"
	"productapp/internal/repository/mysql"
	"productapp/internal/repository/mysql/mysqlproduct"
	"productapp/internal/repository/redis"
	"productapp/internal/repository/redis/productcache"
	productservice "productapp/internal/service"
	productvalidator "productapp/internal/validator"
	"syscall"
	"time"

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

	mgr := migrator.New(cfg2.Mysql)
	mgr.Up()

	MysqlRepo := mysql.New(cfg2.Mysql)
	productRepo := mysqlproduct.New(MysqlRepo)

	redisAdapter := redis.NewAdapter(cfg2.Redis)
	ctx := context.Background()

	productTTL := time.Duration(cfg2.Redis.ProductTTLMinutes) * time.Minute
	productCache := productcache.NewProductCache(redisAdapter, productTTL)

	if err := redisAdapter.Ping(ctx); err != nil {
		logger.Error("redis unavailable, running without cache: %v", zap.Error(err))
		productCache = nil
	}

	productSvc := productservice.New(productRepo, productCache)
	productV := productvalidator.New()
	grpcServer := grpc.NewServer(productV, *productSvc, cfg2.GrpcServer.ProductAddress, cfg2.Metrics.Port)

	go func() {
		logger.Info("🚀gRPC server started on ",
			zap.String("address", cfg2.GrpcServer.ProductAddress))

		if err := grpcServer.Run(); err != nil {
			logger.Fatal("cannot start grpc server", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	logger.Info("Received shutdown signal. Initiating graceful shutdown...")

	grpcServer.GracefulStop()

	if err := MysqlRepo.Conn().Close(); err != nil {
		logger.Error("failed to close MysqlRepo connection", zap.Error(err))
	}

	if err := redisAdapter.Client().Close(); err != nil {
		logger.Error("failed to close Redis Adapter connection", zap.Error(err))
	}

	logger.Info("Graceful shutdown completed successfully. 🛑")

}
