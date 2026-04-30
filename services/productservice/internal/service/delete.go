package productservice

import (
	"context"
	"myapp/pkg/errmsg"
	"myapp/pkg/richerror"
	"productapp/internal/param"
)

func (s *Service) Delete(ctx context.Context, req param.DeleteProductRequest) error {
	const op = "ProductService.Delete"

	p, err := s.repo.GetByID(ctx, req.ID)
	if err != nil {
		return richerror.New(op).WithErr(err)
	}

	if p.StoreID != req.StoreID {
		return richerror.New(op).WithKind(richerror.KindUnexpected).WithMessage(errmsg.ErrorMsgUserNotAllowed)
	}

	return s.repo.Delete(ctx, req.ID)
}
