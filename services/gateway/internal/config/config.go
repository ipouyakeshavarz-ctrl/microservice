package config

import (
	"time"
)

type Config struct {
	HttpServer  Httpserver  `koanf:"http_server"`
	Application Application `koanf:"application"`
	GrpcClient  GrpcClient  `koanf:"grpc_client"`
}

type Application struct {
	GracefulShutdownTimeout time.Duration `koanf:"graceful_shutdown_timeout"`
}
type GrpcClient struct {
	ProductAddress string `koanf:"product_address"`
	UserAddress    string `koanf:"user_address"`
	StoreAddress   string `koanf:"store_address"`
	AuthAddress    string `koanf:"auth_address"`
}

type Httpserver struct {
	Address string `koanf:"address"`
}
