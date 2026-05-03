package storeservice

import (
	"context"
	"storeapp/internal/domain"
	"storeapp/internal/param"
)

type Repository interface {
	CreateStore(ctx context.Context, s domain.Store) (*domain.Store, error)
	UpdateStore(ctx context.Context, s domain.Store) (*domain.Store, error)
	DeleteStore(ctx context.Context, id uint) error
	GetStoreByID(ctx context.Context, id uint) (*domain.Store, error)
	ListStoresByUser(ctx context.Context, userID uint) ([]param.StoreInfo, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{repo: repo}
}
