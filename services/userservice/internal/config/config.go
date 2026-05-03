package config

import (
	"time"
	"userapp/internal/repository/mysql"
)

type Config struct {
	Application Application  `koanf:"application"`
	Grpcserver  Grpcserver   `koanf:"http_server"`
	Mysql       mysql.Config `koanf:"mysql"`
}

type Application struct {
	GracefulShutdownTimeout time.Duration `koanf:"graceful_shutdown_timeout"`
}
type Grpcserver struct {
	AuthPort int `koanf:"auth_port"`
	UserPort int `koanf:"user_port"`
}
