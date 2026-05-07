package productservice

import (
	"context"
	"myapp/pkg/richerror"
	"productapp/internal/param"
)

func (s *Service) GetByID(ctx context.Context, req param.GetProductByIDRequest) (param.GetProductResponse, error) {
	const op = "ProductService.GetByID"
	if s.productCache != nil {
		p, err := s.productCache.GetByID(ctx, req.ID)
		if err != nil {
			return param.GetProductResponse{}, err
		}
		if p != nil {
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
