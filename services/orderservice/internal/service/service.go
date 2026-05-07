package orderservice

import (
	"context"
	"orderapp/internal/domain"
)

type OrderRepository interface {
	Create(ctx context.Context, order *domain.Order) error
	ExistsByCheckoutID(ctx context.Context, checkoutID string) (bool, error)
}

type Service struct {
	repo OrderRepository
}

func New(repo OrderRepository) *Service {
	return &Service{
		repo: repo,
	}
}
