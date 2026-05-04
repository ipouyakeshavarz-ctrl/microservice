package config

import (
	"storeapp/internal/repository/mysql"
	"storeapp/internal/repository/redis"
	"time"
)

type Config struct {
	Application Application  `koanf:"application"`
	GrpcServer  GrpcServer   `koanf:"grpc_server"`
	Mysql       mysql.Config `koanf:"mysql"`
	Redis       redis.Config `koanf:"redis"`
}
type Application struct {
	GracefulShutdownTimeout time.Duration `koanf:"graceful_shutdown_timeout"`
}
type GrpcServer struct {
	StoreAddress string `koanf:"store_address"`
}
