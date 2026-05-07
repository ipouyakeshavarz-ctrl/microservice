package carthandler

import (
	"cartapp/internal/param"
	"context"
	"myapp/api/gen/cart"
	"myapp/pkg/richerror"
)

func (h *Handler) Checkout(ctx context.Context,
	req *cart.CheckoutRequest) (*cart.CheckoutResponse, error) {
	const op = "carthandler.Checkout"

	checkoutID, err := h.cartSvc.Checkout(ctx, param.CheckOutRequest{
		UserID: uint(req.UserId),
	})
	if err != nil {
		return nil, richerror.New(op).WithErr(err).WithFields(map[string]string{
			"massage": err.Error(),
		})
	}
	return &cart.CheckoutResponse{
		CheckoutId: checkoutID.CheckoutID,
	}, nil
}
