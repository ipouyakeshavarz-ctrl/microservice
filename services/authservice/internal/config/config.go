package config

import authservice "authapp/internal/service"

type Config struct {
	Auth authservice.Config `koanf:"auth"`
}
