package config

import (
	"orderapp/internal/repository/mysql"
	"time"
)

type Config struct {
	Application Application  `koanf:"application"`
	GrpcServer  GrpcServer   `koanf:"grpc_server"`
	Mysql       mysql.Config `koanf:"mysql"`
	Logger      Logger       `koanf:"logger"`
	RabbitMQ    RabbitMQ     `koanf:"rabbitmq"`
	Metrics     Metrics      `koanf:"metrics"`
}

type Application struct {
	GracefulShutdownTimeout time.Duration `koanf:"graceful_shutdown_timeout"`
}
type GrpcServer struct {
	OrderAddress string `koanf:"order_address"`
}

type Logger struct {
	Development bool   `koanf:"development"`
	ServiceName string `koanf:"service_name"`
	FilePath    string `koanf:"file_path"`
}

type RabbitMQ struct {
	URL           string `koanf:"url"`
	Exchange      string `koanf:"exchange"`
	CheckoutQueue string `koanf:"checkout_queue"`
}

type Metrics struct {
	Port int `koanf:"port"`
}
