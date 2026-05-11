package config

import (
	"cartapp/internal/adapter/redis"
	"time"
)

type Config struct {
	Application Application  `koanf:"application"`
	GrpcServer  GrpcServer   `koanf:"grpc_server"`
	Redis       redis.Config `koanf:"redis"`
	Logger      Logger       `koanf:"logger"`
	RabbitMQ    RabbitMQ     `koanf:"rabbitmq"`
	Metrics     Metrics      `koanf:"metrics"`
}

type Application struct {
	GracefulShutdownTimeout time.Duration `koanf:"graceful_shutdown_timeout"`
}
type GrpcServer struct {
	CartAddress string `koanf:"cart_address"`
}

type Logger struct {
	Development bool   `koanf:"development"`
	ServiceName string `koanf:"service_name"`
	FilePath    string `koanf:"file_path"`
}

type RabbitMQ struct {
	URL                string `koanf:"url"`
	Exchange           string `koanf:"exchange"`
	CheckoutRoutingKey string `koanf:"checkout_routing_key"`
}

type Metrics struct {
	Port int `koanf:"port"`
}
