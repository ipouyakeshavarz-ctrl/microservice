package carthandler

import (
	"cartapp/internal/param"
	"context"
	"myapp/api/gen/cart"
	"myapp/pkg/richerror"
)

func (h *Handler) GetCart(ctx context.Context,
	req *cart.GetCartRequest) (*cart.GetCartResponse, error) {
	const op = "carthandler.GetCart"

	resp, err := h.cartSvc.GetCart(ctx, param.GetCartRequest{
		UserID: uint(req.UserId),
	})
	if err != nil {
		return nil, richerror.New(op).WithErr(err).WithFields(map[string]string{
			"massage": err.Error(),
		})
	}
	items := make([]*cart.CartItem, 0, len(resp.Items))

	for _, item := range resp.Items {
		items = append(items, &cart.CartItem{
			ProductId: uint64(item.ProductID),
			Quantity:  int32(item.Quantity),
		})
	}

	return &cart.GetCartResponse{
		Cart: &cart.Cart{
			UserId: uint64(resp.UserID),
			Items:  items,
		},
	}, nil
}
