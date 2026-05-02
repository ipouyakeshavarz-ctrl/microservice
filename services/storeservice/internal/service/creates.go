package storeservice

import (
	"context"
	"myapp/pkg/richerror"
	"storeapp/internal/entity"
	"storeapp/internal/param"
)

// todo validitor
func (s *Service) CreateStore(ctx context.Context, req param.CreateStoreRequest) (param.CreateStoreResponse, error) {
	const op = "StoreService.CreateStore"

	store := &entity.Store{
		UserID:      req.UserID,
		Name:        req.Name,
		Description: req.Description,
		PhoneNumber: req.PhoneNumber,
		LogoURL:     req.LogoURL,
		//Address: domain.Address{
		//	Street:   req.Address.Street,
		//	City:     req.Address.City,
		//	Province: req.Address.Province,
		//},
		IsActive: true,
	}
	finalStore, err := s.repo.CreateStore(ctx, *store)

	if err != nil {
		return param.CreateStoreResponse{}, richerror.New(op).WithErr(err)
	}

	return param.CreateStoreResponse{Store: param.StoreInfo{
		ID:          finalStore.ID,
		Name:        finalStore.Name,
		Description: finalStore.Description,
		PhoneNumber: finalStore.PhoneNumber,
		LogoURL:     finalStore.LogoURL,
		Address: entity.Address{
			Street:      store.Address.Street,
			City:        store.Address.City,
			Description: store.Address.Description,
		},
	}}, nil
}
