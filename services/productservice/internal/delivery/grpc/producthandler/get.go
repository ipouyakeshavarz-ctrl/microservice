package producthandler

import (
	"context"
	"myapp/api/gen/product"
	"myapp/pkg/richerror"
	"productapp/internal/param"
)

func (h *Handler) GetProductByID(ctx context.Context,
	req *product.GetProductByIDRequest) (*product.GetProductResponse, error) {
	const op = "productHandler.GetProductByID"

	input := param.GetProductByIDRequest{
		ID:      uint(req.ProductId),
		StoreID: uint(req.StoreId),
	}

	resp, err := h.productSvc.GetByID(ctx, input)
	if err != nil {
		return nil, richerror.New(op).WithErr(err)
	}

	return &product.GetProductResponse{
		Product: &product.ProductInfo{
			Id:          uint64(resp.Product.ID),
			StoreId:     uint64(resp.Product.StoreID),
			Name:        resp.Product.Name,
			Description: resp.Product.Description,
			Category:    string(resp.Product.Category),
			Price:       float32(resp.Product.Price),
			Stock:       int64(resp.Product.Stock),
			Sku:         resp.Product.SKU,
			ImageUrl:    resp.Product.ImageURL,
			IsActive:    resp.Product.IsActive,
		}}, nil
}
