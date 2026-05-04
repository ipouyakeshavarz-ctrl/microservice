package storecache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"myapp/pkg/richerror"
	"storeapp/internal/domain"
	redisadapter "storeapp/internal/repository/redis"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	defaultProductTTL = 10 * time.Minute
)

type StoreCache struct {
	adapter *redisadapter.Adapter
	ttl     time.Duration
}

func NewStoreCache(adapter *redisadapter.Adapter, ttl time.Duration) *StoreCache {
	if ttl <= 0 {
		ttl = defaultProductTTL
	}
	return &StoreCache{adapter: adapter, ttl: ttl}
}

func storeKey(id uint) string {
	return fmt.Sprintf("store:id:%d", id)
}

func (c *StoreCache) GetByID(ctx context.Context, id uint) (*domain.Store, error) {
	const op = "store.cache.GetByID"
	key := storeKey(id)

	val, err := c.adapter.Client().Get(ctx, key).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, nil
		}
		return nil, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}

	var p domain.Store
	if err := json.Unmarshal([]byte(val), &p); err != nil {
		return nil, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}

	return &p, nil
}

func (c *StoreCache) SetByID(ctx context.Context, id uint, p *domain.Store) error {
	const op = "store.cache.SetByID"
	key := storeKey(id)

	b, err := json.Marshal(p)
	if err != nil {
		return richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}

	return c.adapter.Client().Set(ctx, key, b, c.ttl).Err()
}

func (c *StoreCache) DeleteByID(ctx context.Context, id uint) error {
	const op = "store.cache.DeleteByID"
	key := storeKey(id)
	err := c.adapter.Client().Del(ctx, key).Err()
	if err != nil {
		return richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}
	return nil
}

func (c *StoreCache) GetManyByIDs(ctx context.Context, ids []uint) (map[uint]*domain.Store, error) {
	const op = "store.cache.GetManyByIDs"

	if len(ids) == 0 {
		return map[uint]*domain.Store{}, nil
	}

	keys := make([]string, len(ids))
	for i, id := range ids {
		keys[i] = storeKey(id)
	}

	values, err := c.adapter.Client().MGet(ctx, keys...).Result()
	if err != nil {
		return nil, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}

	out := make(map[uint]*domain.Store)

	for i, raw := range values {
		if raw == nil {
			continue
		}

		id := ids[i]

		var p domain.Store
		switch v := raw.(type) {
		case string:
			if err := json.Unmarshal([]byte(v), &p); err != nil {
				return nil, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
			}
		case []byte:
			if err := json.Unmarshal(v, &p); err != nil {
				return nil, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
			}
		default:
			return nil, richerror.New(op).WithKind(richerror.KindUnexpected)
		}

		out[id] = &p
	}

	return out, nil
}
