package storeservice

import (
	"context"
	"myapp/pkg/richerror"
	"storeapp/internal/param"
)

func (s *Service) DeleteStore(ctx context.Context, req param.DeleteStoreRequest) error {
	const op = "StoreService.DeleteStore"
	_, err := s.repo.GetStoreByID(ctx, req.StoreID)
	if err != nil {
		return richerror.New(op).WithErr(err)
	}
	err = s.repo.DeleteStore(ctx, req.StoreID)
	return richerror.New(op).WithErr(err)
}
