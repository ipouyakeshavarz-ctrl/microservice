package config

import (
	"productapp/internal/repository/mysql"
	"time"
)

type Config struct {
	Application Application  `koanf:"application"`
	Httpserver  Httpserver   `koanf:"http_server"`
	Mysql       mysql.Config `koanf:"mysql"`
}

type Application struct {
	GracefulShutdownTimeout time.Duration `koanf:"graceful_shutdown_timeout"`
}
type Httpserver struct {
	Port int `koanf:"port"`
}
