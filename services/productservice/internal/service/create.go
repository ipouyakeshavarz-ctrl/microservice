package productservice

import (
	"context"
	"myapp/pkg/errmsg"
	"myapp/pkg/richerror"
	"productapp/internal/domain"
	"productapp/internal/param"
)

func (s *Service) Create(ctx context.Context, p param.CreateProductRequest) (param.CreateProductResponse, error) {
	const op = "ProductService.Create"
	if !p.Category.IsValid() {
		return param.CreateProductResponse{}, richerror.New(op).WithKind(richerror.KindInvalid).WithMessage(errmsg.ErrorMsgCategoryIsNotValid)
	}

	exists, err := s.repo.SKUExists(ctx, p.SKU)
	if err != nil {
		return param.CreateProductResponse{}, err
	}

	if exists {
		return param.CreateProductResponse{}, richerror.New(op).
			WithKind(richerror.KindInvalid).
			WithMessage(errmsg.ErrorMsgSkuAlreadyExists)
	}

	createdProduct, err := s.repo.Create(ctx, domain.Product{
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
		return param.CreateProductResponse{}, richerror.New(op).
			WithKind(richerror.KindUnexpected).
			WithMessage(errmsg.ErrorMsgFailedToCreateProductInDB).
			WithErr(err)
	}

	return param.CreateProductResponse{Product: param.ProductInfo{
		StoreID:     createdProduct.StoreID,
		ID:          createdProduct.ID,
		Name:        createdProduct.Name,
		Description: createdProduct.Description,
		Category:    createdProduct.Category,
		Price:       createdProduct.Price,
		Stock:       createdProduct.Stock,
		SKU:         createdProduct.SKU,
		ImageURL:    createdProduct.ImageURL,
		IsActive:    createdProduct.IsActive,
	}}, nil
}
