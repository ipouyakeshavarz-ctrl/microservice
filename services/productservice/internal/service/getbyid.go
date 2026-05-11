package productservice

import (
	"context"
	"myapp/pkg/richerror"
	"productapp/internal/param"
)

func (s *Service) GetByID(ctx context.Context, req param.GetProductByIDRequest) (param.GetProductResponse, error) {
	const op = "ProductService.GetByID"

	if s.productCache != nil {

		cached, err := s.productCache.GetByID(ctx, req.ID)

		if err == nil && cached != nil {

			return param.GetProductResponse{Product: param.ProductInfo{
				ID:          cached.ID,
				StoreID:     cached.StoreID,
				Name:        cached.Name,
				Description: cached.Description,
				Category:    cached.Category,
				Price:       cached.Price,
				Stock:       cached.Stock,
				SKU:         cached.SKU,
				ImageURL:    cached.ImageURL,
				IsActive:    cached.IsActive,
			}}, nil
		}

	}

	p, err := s.repo.GetByID(ctx, req.ID)

	if err != nil {

		if re, ok := err.(*richerror.RichError); ok {
			return param.GetProductResponse{}, re
		}

		return param.GetProductResponse{}, richerror.New(op).
			WithKind(richerror.KindUnexpected).
			WithErr(err)
	}

	if s.productCache != nil && p != nil {
		_ = s.productCache.SetByID(ctx, req.ID, p)
	}

	return param.GetProductResponse{Product: param.ProductInfo{
		ID:          p.ID,
		StoreID:     p.StoreID,
		Name:        p.Name,
		Description: p.Description,
		Category:    p.Category,
		Price:       p.Price,
		Stock:       p.Stock,
		SKU:         p.SKU,
		ImageURL:    p.ImageURL,
		IsActive:    p.IsActive,
	}}, nil
}
