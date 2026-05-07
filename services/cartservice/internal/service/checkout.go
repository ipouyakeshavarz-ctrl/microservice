package cartservice

import (
	"cartapp/internal/domain"
	"cartapp/internal/param"
	"context"
	"fmt"
	"myapp/pkg/errmsg"
	"myapp/pkg/richerror"
	"time"
)

func (s *Service) Checkout(ctx context.Context, req param.CheckOutRequest) (*param.CheckoutResponse, error) {
	const op = "cartservice.Checkout"

	cart, err := s.repo.Get(ctx, req.UserID)
	if err != nil {
		return &param.CheckoutResponse{}, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}
	if cart == nil {
		return &param.CheckoutResponse{}, richerror.New(op).
			WithMessage(errmsg.ErrorMsgCartNotFound).WithKind(richerror.KindNotFound)
	}
	if cart.IsEmpty() {
		return &param.CheckoutResponse{}, richerror.New(op).
			WithMessage(errmsg.ErrorMsgCartIsEmpty).WithKind(richerror.KindNotFound)
	}

	checkoutID := s.idgen()

	items := make([]domain.CheckedOutItem, 0, len(cart.Items))
	for _, item := range cart.Items {
		items = append(items, domain.CheckedOutItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		})
	}

	event := domain.CartCheckedOutEvent{
		CheckoutID: checkoutID,
		UserID:     cart.UserID,
		Items:      items,
		OccurredAt: time.Now(),
	}

	fmt.Println(event)

	if err := s.publisher.PublishCartCheckedOut(ctx, event); err != nil {
		return &param.CheckoutResponse{}, err
	}

	if err := s.repo.Delete(ctx, req.UserID); err != nil {
		return &param.CheckoutResponse{}, err
	}

	return &param.CheckoutResponse{
		CheckoutID: checkoutID,
	}, nil
}
