package config

import "os"

type Config struct {
	Port               string
	AuthServiceAddr    string
	UserServiceAddr    string
	ProductServiceAddr string
}

func Load() *Config {

	return &Config{
		Port:               getEnv("PORT", ":8080"),
		AuthServiceAddr:    getEnv("AUTH_SERVICE", "localhost:5001"),
		UserServiceAddr:    getEnv("USER_SERVICE", "localhost:5002"),
		ProductServiceAddr: getEnv("PRODUCT_SERVICE", "localhost:5003"),
	}
}

func getEnv(key, fallback string) string {

	v := os.Getenv(key)

	if v == "" {
		return fallback
	}

	return v
}
