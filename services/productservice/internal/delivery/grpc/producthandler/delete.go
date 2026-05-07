package producthandler

import (
	"context"
	"myapp/api/gen/product"
	"myapp/pkg/richerror"
	"productapp/internal/param"
)

func (h *Handler) DeleteProduct(ctx context.Context,
	req *product.DeleteProductRequest) (*product.DeleteProductResponse, error) {
	const op = "productHandler.DeleteProduct"

	input := param.DeleteProductRequest{
		ID:      uint(req.ProductId),
		StoreID: uint(req.StoreId),
		UserID:  uint(req.UserId),
	}

	resp, err := h.productSvc.Delete(ctx, input)
	if err != nil {
		return nil, richerror.New(op).WithErr(err).WithFields(map[string]string{
			"massage": err.Error(),
		})
	}
	return &product.DeleteProductResponse{
		Success: resp.Success,
	}, nil
}
