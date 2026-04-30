package config

import (
	"time"
	"userapp/internal/auth"
	"userapp/internal/repository/mysql"
)

type Config struct {
	Application Application        `koanf:"application"`
	Httpserver  Httpserver         `koanf:"http_server"`
	Auth        authservice.Config `koanf:"auth"`
	Mysql       mysql.Config       `koanf:"mysql"`
}

type Application struct {
	GracefulShutdownTimeout time.Duration `koanf:"graceful_shutdown_timeout"`
}
type Httpserver struct {
	Port int `koanf:"port"`
}
