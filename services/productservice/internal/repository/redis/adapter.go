package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type Config struct {
	Host              string `koanf:"host"`
	Port              int    `koanf:"port"`
	Password          string `koanf:"password"`
	DB                int    `koanf:"db"`
	ProductTTLMinutes int    `koanf:"product_ttl_minutes"`
}

type Adapter struct {
	client *redis.Client
}

func NewAdapter(cfg Config) *Adapter {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	})
	return &Adapter{client: rdb}
}

func (a *Adapter) Client() *redis.Client {
	return a.client
}

func (a *Adapter) Ping(ctx context.Context) error {
	return a.client.Ping(ctx).Err()
}
