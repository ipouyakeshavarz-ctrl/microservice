package producthandler

import (
	"context"
	"myapp/api/gen/product"
	"myapp/pkg/errmsg"
	"myapp/pkg/richerror"
	"productapp/internal/entity"

	"productapp/internal/param"
)

func (h *Handler) CreateProduct(ctx context.Context,
	req *product.CreateProductRequest) (*product.CreateProductResponse, error) {
	const op = "productHandler.CreateProduct"
	category := entity.Category(req.GetCategory())
	ok := category.IsValid()
	if !ok {
		return &product.CreateProductResponse{}, richerror.New(op).
			WithMessage(errmsg.ErrorMsgCategoryIsNotValid).WithKind(richerror.KindNotFound)
	}

	input := param.CreateProductRequest{
		StoreID:     uint(req.StoreId),
		Name:        req.Name,
		Description: req.Description,
		Category:    category,
		Price:       float64(req.Price),
		Stock:       int(req.Stock),
		SKU:         req.Sku,
		ImageURL:    req.ImageUrl,
		IsActive:    req.IsActive,
	}

	resp, err := h.productSvc.Create(ctx, input)
	if err != nil {
		return nil, err
	}

	return &product.CreateProductResponse{
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
		},
	}, nil
}
