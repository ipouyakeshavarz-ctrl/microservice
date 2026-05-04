package main

import (
	"context"
	"fmt"
	"log"
	"myapp/pkg/config"
	cfg "productapp/internal/config"
	"productapp/internal/delivery/grpc"
	"productapp/internal/repository/migrator"
	"productapp/internal/repository/mysql"
	"productapp/internal/repository/mysql/mysqlproduct"
	"productapp/internal/repository/redis"
	"productapp/internal/repository/redis/productcache"
	productservice "productapp/internal/service"
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
	productRepo := mysqlproduct.New(MysqlRepo)

	redisAdapter := redis.NewAdapter(cfg2.Redis)
	ctx := context.Background()

	productTTL := time.Duration(cfg2.Redis.ProductTTLMinutes) * time.Minute
	productCache := productcache.NewProductCache(redisAdapter, productTTL)

	if err := redisAdapter.Ping(ctx); err != nil {
		log.Printf("redis unavailable, running without cache: %v", err)
		productCache = nil
	}

	productSvc := productservice.New(productRepo, productCache)

	grpcServer := grpc.NewServer(*productSvc, cfg2.GrpcServer.ProductAddress)

	if err := grpcServer.Run(); err != nil {
		log.Fatal(err)
	}
}
