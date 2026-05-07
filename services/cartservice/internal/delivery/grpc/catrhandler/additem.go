package carthandler

import (
	"cartapp/internal/param"
	"context"
	"myapp/api/gen/cart"
	"myapp/pkg/richerror"
)

func (h *Handler) AddItem(ctx context.Context,
	req *cart.AddItemRequest) (*cart.AddItemResponse, error) {
	const op = "carthandler.AddItem"

	err := h.cartSvc.AddItem(ctx, param.AddItemRequest{
		UserID:    uint(req.UserId),
		ProductID: uint(req.ProductId),
		Quantity:  uint(req.Quantity),
	})
	if err != nil {
		return nil, richerror.New(op).WithErr(err).WithFields(map[string]string{
			"massage": err.Error(),
		})
	}
	return &cart.AddItemResponse{}, nil
}
