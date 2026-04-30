package storeservice

import (
	"context"
	"storeapp/internal/entity"
)

type Repository interface {
	CreateStore(ctx context.Context, s entity.Store) (*entity.Store, error)
	UpdateStore(ctx context.Context, s entity.Store) (*entity.Store, error)
	DeleteStore(ctx context.Context, id uint) error
	GetStoreByID(ctx context.Context, id uint) (*entity.Store, error)
	ListStoresByUser(ctx context.Context, userID uint) ([]entity.Store, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{repo: repo}
}
