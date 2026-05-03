package config

import authservice "authapp/internal/service"

type Config struct {
	Auth       authservice.Config `koanf:"auth"`
	GrpcServer GrpcServer         `koanf:"grpc_server"`
}

type GrpcServer struct {
	AuthAddress string `koanf:"auth_address"`
}
