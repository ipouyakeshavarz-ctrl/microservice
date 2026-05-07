package cartservice

import (
	"cartapp/internal/domain"
	"cartapp/internal/param"
	"context"
	"myapp/pkg/richerror"
	"time"
)

func (s *Service) GetCart(ctx context.Context, req param.GetCartRequest) (*param.GetCartResponse, error) {
	const op = "CartService.GetCart"
	cart, err := s.repo.Get(ctx, req.UserID)
	if err != nil {
		return &param.GetCartResponse{}, richerror.New(op).WithKind(richerror.KindUnexpected).WithErr(err)
	}

	if cart == nil {
		cart = domain.NewCart(req.UserID, time.Now())
	}

	items := make([]param.CartItemView, 0, len(cart.Items))
	for _, i := range cart.Items {
		items = append(items, param.CartItemView{
			ProductID: i.ProductID,
			Quantity:  i.Quantity,
		})
	}

	return &param.GetCartResponse{
		UserID: cart.UserID,
		Items:  items,
	}, nil
}
