package productservice

import (
	"context"
	"productapp/internal/domain"
)

type Repository interface {
	Create(ctx context.Context, p domain.Product) (*domain.Product, error)
	Update(ctx context.Context, p domain.Product) (*domain.Product, error)
	Delete(ctx context.Context, id uint) error
	GetByID(ctx context.Context, id uint) (*domain.Product, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{repo: repo}
}
