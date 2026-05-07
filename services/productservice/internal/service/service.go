package productservice

import (
	"context"
	"productapp/internal/domain"
	"productapp/internal/repository/redis/productcache"
)

type Repository interface {
	SKUExists(ctx context.Context, sku string) (bool, error)
	Create(ctx context.Context, p domain.Product) (*domain.Product, error)
	Update(ctx context.Context, p domain.Product) (*domain.Product, error)
	Delete(ctx context.Context, id uint) error
	GetByID(ctx context.Context, id uint) (*domain.Product, error)
}

type Service struct {
	repo         Repository
	productCache *productcache.ProductCache
}

func New(repo Repository, productCache *productcache.ProductCache) *Service {
	return &Service{repo: repo, productCache: productCache}
}
