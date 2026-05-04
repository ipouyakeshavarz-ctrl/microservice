package productcache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"myapp/pkg/richerror"
	"productapp/internal/domain"
	redisadapter "productapp/internal/repository/redis"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	defaultProductTTL = 10 * time.Minute
)

type ProductCache struct {
	adapter *redisadapter.Adapter
	ttl     time.Duration
}

func NewProductCache(adapter *redisadapter.Adapter, ttl time.Duration) *ProductCache {
	if ttl <= 0 {
		ttl = defaultProductTTL
	}
	return &ProductCache{adapter: adapter, ttl: ttl}
}

func productKey(id uint) string {
	return fmt.Sprintf("product:id:%d", id)
}

func (c *ProductCache) GetByID(ctx context.Context, id uint) (*domain.Product, error) {
	const op = "product.cache.GetByID"
	key := productKey(id)

	val, err := c.adapter.Client().Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}
		return nil, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}

	var p domain.Product
	if err := json.Unmarshal([]byte(val), &p); err != nil {
		return nil, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}

	return &p, nil
}

func (c *ProductCache) SetByID(ctx context.Context, id uint, p *domain.Product) error {
	const op = "product.cache.SetByID"
	key := productKey(id)

	b, err := json.Marshal(p)
	if err != nil {
		return richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}

	return c.adapter.Client().Set(ctx, key, b, c.ttl).Err()
}

func (c *ProductCache) DeleteByID(ctx context.Context, id uint) error {
	const op = "product.cache.DeleteByID"
	key := productKey(id)
	err := c.adapter.Client().Del(ctx, key).Err()
	if err != nil {
		return richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}
	return nil
}
