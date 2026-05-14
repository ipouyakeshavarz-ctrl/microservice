package producthandler

import (
	"context"
	"myapp/api/gen/product"
	"myapp/pkg/errmsg"
	"myapp/pkg/richerror"
	"productapp/internal/domain"
	"productapp/internal/param"
)

func (h *Handler) UpdateProduct(ctx context.Context, req *product.UpdateProductRequest) (*product.UpdateProductResponse, error) {
	const op = "producthandler.UpdateProduct"

	var category domain.Category

	if req.Category != "" {
		category = domain.Category(req.GetCategory())
		ok := category.IsValid()
		if !ok {
			return nil, richerror.New(op).
				WithKind(richerror.KindInvalid).
				WithMessage(errmsg.ErrorMsgCategoryIsNotValid)
		}
	}
	input := param.UpdateProductRequest{
		ID:          uint(req.Id),
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

	if fieldErrors, err := h.productV.ValidateUpdateRequest(input); err != nil {
		return nil, richerror.New(op).
			WithKind(richerror.KindInvalid).
			WithMessage(errmsg.ErrorMsgInvalidInput).
			WithFields(fieldErrors)
	}

	resp, err := h.productSvc.Update(ctx, input)

	if err != nil {

		return nil, richerror.New(op).WithErr(err).WithFields(map[string]string{
			"massage": err.Error(),
		})
	}

	return &product.UpdateProductResponse{
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
