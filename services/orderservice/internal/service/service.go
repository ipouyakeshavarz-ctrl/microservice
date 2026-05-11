package orderservice

import (
	"context"
	"orderapp/internal/domain"
)

type OrderRepository interface {
	Create(ctx context.Context, order *domain.Order) error
	ExistsByCheckoutID(ctx context.Context, checkoutID string) (bool, error)
	GetByID(ctx context.Context, id uint) (*domain.Order, error)
	ListByUserID(ctx context.Context, userID uint) ([]domain.Order, error)
}

type Service struct {
	repo OrderRepository
}

func New(repo OrderRepository) *Service {
	return &Service{
		repo: repo,
	}
}
