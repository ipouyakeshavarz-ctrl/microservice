package main

import (
	"context"
	"fmt"
	"log"
	"myapp/pkg/config"
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
)

func main() {
	var cfg2 cfg.Config
	err := config.Load("config.yml", &cfg2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("cfg:%v\n", cfg2)

	mgr := migrator.New(cfg2.Mysql)
	mgr.Up()

	MysqlRepo := mysql.New(cfg2.Mysql)
	storeRepo := mysqlstore.New(MysqlRepo)

	redisAdapter := redis.NewAdapter(cfg2.Redis)
	ctx := context.Background()

	storeTTL := time.Duration(cfg2.Redis.StoreTTLMinutes) * time.Minute
	storeCache := storecache.NewStoreCache(redisAdapter, storeTTL)

	if err := redisAdapter.Ping(ctx); err != nil {
		log.Printf("redis unavailable, running without cache: %v", err)
		storeCache = nil
	}

	storeSvc := storeservice.New(storeRepo, storeCache)
	storeV := storevalidator.New()

	grpcServer := grpc.NewServer(storeV, storeSvc, cfg2.GrpcServer.StoreAddress)

	if err := grpcServer.Run(); err != nil {
		log.Fatal(err)
	}
}
