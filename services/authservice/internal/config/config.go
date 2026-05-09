package config

import authservice "authapp/internal/service"

type Config struct {
	Auth       authservice.Config `koanf:"auth"`
	GrpcServer GrpcServer         `koanf:"grpc_server"`
	Logger     Logger             `koanf:"logger"`
	Metrics    Metrics            `koanf:"metrics"`
}

type GrpcServer struct {
	AuthAddress string `koanf:"auth_address"`
}

type Logger struct {
	Development bool   `koanf:"development"`
	ServiceName string `koanf:"service_name"`
	FilePath    string `koanf:"file_path"`
}

type Metrics struct {
	Port int `koanf:"port"`
}
