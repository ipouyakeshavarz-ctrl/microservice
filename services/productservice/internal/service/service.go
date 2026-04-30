package productservice

import (
	"context"
	"productapp/internal/entity"
)

type Repository interface {
	Create(ctx context.Context, p entity.Product) (*entity.Product, error)
	Update(ctx context.Context, p entity.Product) (*entity.Product, error)
	Delete(ctx context.Context, id uint) error
	GetByID(ctx context.Context, id uint) (*entity.Product, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{repo: repo}
}
