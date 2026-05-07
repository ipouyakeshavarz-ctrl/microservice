package cartservice

import (
	"cartapp/internal/domain"
	"cartapp/internal/param"
	"context"
	"myapp/pkg/errmsg"
	"myapp/pkg/richerror"
	"time"
)

func (s *Service) AddItem(ctx context.Context, req param.AddItemRequest) error {
	const op = "CartService.AddItem"
	if req.UserID == 0 || req.ProductID == 0 {
		return richerror.New(op).
			WithKind(richerror.KindInvalid).WithMessage(errmsg.ErrorMsgInvalidInput)
	}
	if req.Quantity == 0 {
		return richerror.New(op).WithKind(richerror.
			KindInvalid).WithMessage(errmsg.ErrorMsgInvalidQty)
	}

	cart, err := s.repo.Get(ctx, req.UserID)
	if err != nil {
		if re, ok := err.(*richerror.RichError); ok {
			return re
		}

		return richerror.New(op).
			WithKind(richerror.KindUnexpected).
			WithMessage(errmsg.ErrorMsgFailedToGetCartInDB).
			WithErr(err)
	}
	if cart == nil {
		cart = domain.NewCart(req.UserID, time.Now())
	}

	cart.AddItem(req.ProductID, req.Quantity, time.Now())

	if err := s.repo.Save(ctx, cart); err != nil {
		if re, ok := err.(*richerror.RichError); ok {
			return re
		}

		return richerror.New(op).
			WithKind(richerror.KindUnexpected).
			WithMessage(errmsg.ErrorMsgFailedToGetCartInDB).
			WithErr(err)
	}

	return nil
}
