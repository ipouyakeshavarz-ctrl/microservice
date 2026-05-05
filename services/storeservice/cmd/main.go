package main

import (
	"context"
	"myapp/pkg/config"
	"myapp/pkg/logger"
	cfg "storeapp/internal/config"
	"storeapp/internal/delivery/grpc"
	"storeapp/internal/repository/migrator"
	"storeapp/internal/repository/mysql"
	"storeapp/internal/repository/mysql/mysqlstore"
	"storeapp/internal/repository/redis"
	"storeapp/internal/repository/redis/storecache"
	storeservice "storeapp/internal/service"
	storevalidator "storeapp/internal/validator"
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
	storeRepo := mysqlstore.New(MysqlRepo)

	redisAdapter := redis.NewAdapter(cfg2.Redis)
	ctx := context.Background()

	storeTTL := time.Duration(cfg2.Redis.StoreTTLMinutes) * time.Minute
	storeCache := storecache.NewStoreCache(redisAdapter, storeTTL)

	if err := redisAdapter.Ping(ctx); err != nil {
		logger.Error("redis unavailable, running without cache: %v", zap.Error(err))
		storeCache = nil
	}

	storeSvc := storeservice.New(storeRepo, storeCache)
	storeV := storevalidator.New()

	grpcServer := grpc.NewServer(storeV, storeSvc, cfg2.GrpcServer.StoreAddress)

	logger.Info("🚀gRPC server started on", zap.String("address", cfg2.GrpcServer.StoreAddress))

	if err := grpcServer.Run(); err != nil {
		logger.Fatal("cannot start grpc server", zap.Error(err))
	}
}
