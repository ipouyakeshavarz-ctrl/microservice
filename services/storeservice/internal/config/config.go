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
	Logger      Logger       `koanf:"logger"`
	Metrics     Metrics      `koanf:"metrics"`
}
type Application struct {
	GracefulShutdownTimeout time.Duration `koanf:"graceful_shutdown_timeout"`
}
type GrpcServer struct {
	StoreAddress string `koanf:"store_address"`
}

type Logger struct {
	Development bool   `koanf:"development"`
	ServiceName string `koanf:"service_name"`
	FilePath    string `koanf:"file_path"`
}

type Metrics struct {
	Port int `koanf:"port"`
}
