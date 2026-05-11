package config

import (
	"time"
	"userapp/internal/repository/mysql"
)

type Config struct {
	Application Application  `koanf:"application"`
	GrpcServer  GrpcServer   `koanf:"grpc_server"`
	Mysql       mysql.Config `koanf:"mysql"`
	Logger      Logger       `koanf:"logger"`
	Metrics     Metrics      `koanf:"metrics"`
}

type Application struct {
	GracefulShutdownTimeout time.Duration `koanf:"graceful_shutdown_timeout"`
}
type GrpcServer struct {
	AuthAddress string `koanf:"auth_address"`
	UserAddress string `koanf:"user_address"`
}

type Logger struct {
	Development bool   `koanf:"development"`
	ServiceName string `koanf:"service_name"`
	FilePath    string `koanf:"file_path"`
}

type Metrics struct {
	Port int `koanf:"port"`
}
