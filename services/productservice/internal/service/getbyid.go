package productservice

import (
	"context"
	"myapp/pkg/errmsg"
	"myapp/pkg/richerror"
	"productapp/internal/param"
)

func (s *Service) GetByID(ctx context.Context, req param.GetProductByIDRequest) (param.GetProductResponse, error) {
	const op = "ProductService.GetByID"

	p, err := s.repo.GetByID(ctx, req.ID)
	if err != nil {
		return param.GetProductResponse{}, richerror.New(op).WithErr(err)
	}

	if p.StoreID != req.StoreID {
		return param.GetProductResponse{}, richerror.New(op).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrorMsgUserNotAllowed)
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
