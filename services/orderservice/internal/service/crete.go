package orderservice

import (
	"context"
	"orderapp/internal/domain"
	"orderapp/internal/param"
	"time"
)

func (s *Service) CreateFromCheckout(ctx context.Context, req param.CreateFromCheckoutRequest) error {

	exists, err := s.repo.ExistsByCheckoutID(ctx, req.Event.CheckoutID)
	if err != nil {
		return err
	}

	if exists {
		return nil
	}

	items := make([]domain.OrderItem, 0, len(req.Event.Items))

	for _, i := range req.Event.Items {
		items = append(items, domain.OrderItem{
			ProductID: i.ProductID,
			Quantity:  i.Quantity,
		})
	}

	order := &domain.Order{
		CheckoutID: req.Event.CheckoutID,
		UserID:     req.Event.UserID,
		Status:     domain.OrderStatusCreated,
		Items:      items,
		CreatedAt:  time.Now(),
	}

	return s.repo.Create(ctx, order)
}
