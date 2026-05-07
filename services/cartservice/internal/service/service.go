package cartservice

import (
	"cartapp/internal/domain"
	"context"
	"time"
)

type Repository interface {
	Save(ctx context.Context, cart *domain.Cart) error
	Get(ctx context.Context, userID uint) (*domain.Cart, error)
	Delete(ctx context.Context, userID uint) error
}

type EventPublisher interface {
	PublishCartCheckedOut(ctx context.Context, event domain.CartCheckedOutEvent) error
}

type Service struct {
	repo      Repository
	publisher EventPublisher
	idgen     func() string
}

func New(repo Repository, publisher EventPublisher) *Service {
	return &Service{repo: repo,
		publisher: publisher,
		idgen: func() string {
			return "checkout-" + time.Now().Format("20060102150405")
		}}
}
