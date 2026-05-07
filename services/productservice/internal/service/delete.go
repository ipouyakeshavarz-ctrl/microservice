package productservice

import (
	"context"
	"myapp/pkg/errmsg"
	"myapp/pkg/richerror"
	"productapp/internal/param"
)

func (s *Service) Delete(ctx context.Context, req param.DeleteProductRequest) (param.DeleteProductResponse, error) {
	const op = "ProductService.Delete"

	p, err := s.repo.GetByID(ctx, req.ID)
	if err != nil {
		return param.DeleteProductResponse{
				Success: false,
			}, richerror.New(op).
				WithKind(richerror.KindUnexpected).
				WithMessage(errmsg.ErrorMsgFailedToGetProductInDB).
				WithErr(err)
	}

	if p.StoreID != req.StoreID {
		return param.DeleteProductResponse{
			Success: false,
		}, richerror.New(op).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrorMsgUserNotAllowed)
	}

	dErr := s.repo.Delete(ctx, req.ID)
	if err != nil {

		if re, ok := dErr.(*richerror.RichError); ok {
			return param.DeleteProductResponse{
				Success: false,
			}, re
		}

		return param.DeleteProductResponse{
				Success: false,
			}, richerror.New(op).
				WithKind(richerror.KindUnexpected).
				WithErr(dErr)
	}

	if s.productCache != nil {
		_ = s.productCache.DeleteByID(ctx, req.ID)
	}

	return param.DeleteProductResponse{
		Success: true,
	}, nil
}
