package storeservice

import (
	"context"
	"myapp/pkg/richerror"
	"storeapp/internal/param"
)

func (s *Service) ListStoresByUser(ctx context.Context,
	req param.ListStoresByUserRequest) (param.ListStoresByUserResponse, error) {
	const op = "StoreService.ListStoresByUser"

	stores, err := s.repo.ListStoresByUser(ctx, req.UserID)
	if err != nil {
		return param.ListStoresByUserResponse{}, richerror.New(op).WithErr(err)
	}

	return param.ListStoresByUserResponse{Stores: stores}, nil
}
