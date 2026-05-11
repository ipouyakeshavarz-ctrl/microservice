package storeservice

import (
	"context"
	"myapp/pkg/errmsg"
	"myapp/pkg/richerror"

	"storeapp/internal/param"
)

func (s *Service) DeleteStore(ctx context.Context, req param.DeleteStoreRequest) (param.DeleteStoreResponse, error) {
	const op = "StoreService.DeleteStore"

	store, err := s.repo.GetStoreByID(ctx, req.StoreID)

	if err != nil {

		if re, ok := err.(*richerror.RichError); ok {
			return param.DeleteStoreResponse{Success: false}, re
		}

		return param.DeleteStoreResponse{Success: false},
			richerror.New(op).
				WithKind(richerror.KindNotFound).
				WithMessage(errmsg.ErrorMsgStoreNotFound).
				WithErr(err)
	}

	if store.UserID != req.UserID {
		return param.DeleteStoreResponse{Success: false},
			richerror.New(op).
				WithKind(richerror.KindForbidden).
				WithMessage(errmsg.ErrorMsgUserNotAllowed)
	}

	err = s.repo.DeleteStore(ctx, req.StoreID)

	if err != nil {

		if re, ok := err.(*richerror.RichError); ok {
			return param.DeleteStoreResponse{Success: false}, re
		}

		return param.DeleteStoreResponse{Success: false},
			richerror.New(op).
				WithKind(richerror.KindUnexpected).
				WithMessage(errmsg.ErrorMsgFailedToDeleteStoreInDB).
				WithErr(err)
	}

	if s.storeCache != nil {
		_ = s.storeCache.DeleteByID(ctx, req.StoreID)
	}

	return param.DeleteStoreResponse{
		Success: true,
	}, nil
}
