package cartrepo

import (
	"cartapp/internal/adapter/redis"
	"cartapp/internal/domain"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	goredis "github.com/redis/go-redis/v9"
)

const defaultCartTTL = 10 * time.Minute

type CartRepository struct {
	adapter *redis.Adapter
	ttl     time.Duration
}

func New(adapter *redis.Adapter, ttl time.Duration) *CartRepository {
	if ttl <= 0 {
		ttl = defaultCartTTL
	}

	return &CartRepository{
		adapter: adapter,
		ttl:     ttl,
	}
}

func (r *CartRepository) cartKey(userID uint) string {
	return fmt.Sprintf("cart:user:%d", userID)
}
func (r *CartRepository) Save(ctx context.Context, cart *domain.Cart) error {
	key := r.cartKey(cart.UserID)

	data, err := json.Marshal(cart)
	if err != nil {
		return err
	}

	return r.adapter.Client().
		Set(ctx, key, data, r.ttl).
		Err()
}
func (r *CartRepository) Get(ctx context.Context, userID uint) (*domain.Cart, error) {
	key := r.cartKey(userID)

	data, err := r.adapter.Client().
		Get(ctx, key).
		Bytes()

	if err != nil {
		if errors.Is(err, goredis.Nil) {
			return nil, nil
		}
		return nil, err
	}

	var cart domain.Cart
	if err := json.Unmarshal(data, &cart); err != nil {
		return nil, err
	}

	return &cart, nil
}
func (r *CartRepository) Delete(ctx context.Context, userID uint) error {
	key := r.cartKey(userID)
	return r.adapter.Client().
		Del(ctx, key).
		Err()
}
