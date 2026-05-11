package orderservice

import (
	"context"
	"myapp/pkg/metrics"
	"myapp/pkg/richerror"
	"orderapp/internal/domain"
	"orderapp/internal/param"
	"time"
)

func (s *Service) CreateFromCheckout(ctx context.Context, req param.CreateFromCheckoutRequest) error {
	const op = "orderservice.CreateFromCheckout"

	exists, err := s.repo.ExistsByCheckoutID(ctx, req.Event.CheckoutID)
	if err != nil {
		return richerror.New(op).WithErr(err)
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

	if err := s.repo.Create(ctx, order); err != nil {
		return richerror.New(op).WithErr(err)
	}

	metrics.OrdersCreated.Inc()

	return nil
}
