package storeservice

import (
	"context"
	"storeapp/internal/domain"
	"storeapp/internal/repository/redis/storecache"
)

type Repository interface {
	CreateStore(ctx context.Context, s domain.Store) (*domain.Store, error)
	UpdateStore(ctx context.Context, s domain.Store) (*domain.Store, error)
	DeleteStore(ctx context.Context, id uint) error
	GetStoreByID(ctx context.Context, id uint) (*domain.Store, error)
	ListStoresByUser(ctx context.Context, userID uint) ([]domain.Store, error)
	ListStoreIDsByUser(ctx context.Context, userID uint) ([]uint, error)
	GetStoresByIDs(ctx context.Context, ids []uint) ([]*domain.Store, error)
}

type Service struct {
	repo       Repository
	storeCache *storecache.StoreCache
}

func New(repo Repository, storeCache *storecache.StoreCache) Service {
	return Service{repo: repo, storeCache: storeCache}
}
