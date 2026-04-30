package storeservice

import (
	"context"
	"myapp/pkg/richerror"
	"storeapp/internal/param"
)

func (s *Service) UpdateStore(ctx context.Context, req param.UpdateStoreRequest) (param.UpdateStoreResponse, error) {
	const op = "StoreService.UpdateStore"

	store, err := s.repo.GetStoreByID(ctx, req.ID)
	if err != nil {
		return param.UpdateStoreResponse{}, richerror.New(op).WithErr(err)
	}

	if req.Name != "" {
		store.Name = req.Name
	}

	if req.Description != "" {
		store.Description = req.Description
	}

	if req.PhoneNumber != "" {
		store.PhoneNumber = req.PhoneNumber
	}

	if req.LogoURL != "" {
		store.LogoURL = req.LogoURL
	}

	if req.Address.Street != "" {
		store.Address.Street = req.Address.Street
	}

	if req.Address.City != "" {
		store.Address.City = req.Address.City
	}

	if req.Address.Province != "" {
		store.Address.Province = req.Address.Province
	}

	if req.Address.Description != "" {
		store.Address.Description = req.Address.Description
	}

	if req.Address.PostalCode != "" {
		store.Address.PostalCode = req.Address.PostalCode
	}

	if req.IsActive != store.IsActive {
		store.IsActive = req.IsActive
	}

	updatedStore, err := s.repo.UpdateStore(ctx, *store)
	if err != nil {
		return param.UpdateStoreResponse{}, richerror.New(op).WithErr(err)
	}

	return param.UpdateStoreResponse{Store: param.StoreInfo{
		Name:        updatedStore.Name,
		Description: updatedStore.Description,
		PhoneNumber: updatedStore.PhoneNumber,
		Address:     updatedStore.Address,
		LogoURL:     updatedStore.LogoURL,
		IsActive:    updatedStore.IsActive,
	}}, nil
}
