package config

import (
	"time"
	"userapp/internal/repository/mysql"
)

type Config struct {
	Application Application  `koanf:"application"`
	GrpcServer  GrpcServer   `koanf:"grpc_server"`
	Mysql       mysql.Config `koanf:"mysql"`
}

type Application struct {
	GracefulShutdownTimeout time.Duration `koanf:"graceful_shutdown_timeout"`
}
type GrpcServer struct {
	AuthAddress string `koanf:"auth_address"`
	UserAddress string `koanf:"user_address"`
}
