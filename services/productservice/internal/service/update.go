package productservice

import (
	"context"
	"myapp/pkg/errmsg"
	"myapp/pkg/richerror"
	"productapp/internal/domain"
	"productapp/internal/param"
)

func (s *Service) Update(ctx context.Context, p param.UpdateProductRequest) (param.UpdateProductResponse, error) {
	const op = "ProductService.Update"
	if !p.Category.IsValid() {
		return param.UpdateProductResponse{}, richerror.New(op).WithKind(richerror.KindInvalid).WithMessage(errmsg.ErrorMsgCategoryIsNotValid)
	}

	existing, err := s.repo.GetByID(ctx, p.ID)
	if err != nil {
		return param.UpdateProductResponse{}, richerror.New(op).WithErr(err)
	}

	if existing.StoreID != p.StoreID {
		return param.UpdateProductResponse{}, richerror.New(op).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrorMsgUserNotAllowed)
	}

	updatedProduct, err := s.repo.Update(ctx, domain.Product{
		StoreID:     p.StoreID,
		Name:        p.Name,
		Description: p.Description,
		Category:    p.Category,
		Price:       p.Price,
		Stock:       p.Stock,
		SKU:         p.SKU,
		ImageURL:    p.ImageURL,
		IsActive:    p.IsActive,
	})

	if err != nil {
		return param.UpdateProductResponse{}, richerror.New(op).WithErr(err)
	}

	if s.productCache != nil {
		_ = s.productCache.DeleteByID(ctx, p.ID)
	}

	return param.UpdateProductResponse{
		Product: param.ProductInfo{
			ID:          updatedProduct.ID,
			StoreID:     updatedProduct.StoreID,
			Name:        updatedProduct.Name,
			Description: updatedProduct.Description,
			Category:    updatedProduct.Category,
			Price:       updatedProduct.Price,
			Stock:       updatedProduct.Stock,
			SKU:         updatedProduct.SKU,
			ImageURL:    updatedProduct.ImageURL,
			IsActive:    updatedProduct.IsActive,
		},
	}, nil
}
