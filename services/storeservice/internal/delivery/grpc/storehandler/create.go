package storehandler

import (
	"context"
	"myapp/api/gen/store"
	"myapp/pkg/richerror"
	"storeapp/internal/domain"

	"storeapp/internal/param"
)

func (h *Handler) CreateStore(ctx context.Context, req *store.CreateStoreRequest) (*store.CreateStoreResponse, error) {
	const op = "storeHandler.CreateStore"

	input := param.CreateStoreRequest{
		UserID:      uint(req.UserId),
		Name:        req.Name,
		Description: req.Description,
		LogoURL:     req.LogoUrl,
		Address: domain.Address{
			Street:      req.Address.Street,
			City:        req.Address.City,
			Province:    req.Address.Province,
			PostalCode:  req.Address.PostalCode,
			Description: req.Address.Description,
		},
		PhoneNumber: req.PhoneNumber,
	}

	if fieldErrors, err := h.storeValidator.ValidateCreateRequest(input); err != nil {
		return &store.CreateStoreResponse{}, richerror.New(op).WithMeta(map[string]interface{}{
			"validationErrors": fieldErrors,
		}).WithErr(err)
	}

	resp, err := h.storeSvc.CreateStore(ctx, input)
	if err != nil {
		return nil, err
	}

	return &store.CreateStoreResponse{Store: &store.StoreInfo{
		Id:          uint64(resp.Store.ID),
		UserId:      uint64(resp.Store.UserID),
		Name:        resp.Store.Name,
		Description: resp.Store.Description,
		LogoUrl:     resp.Store.LogoURL,
		Address: &store.Address{Street: resp.Store.Address.Street,
			City:        resp.Store.Address.City,
			Province:    resp.Store.Address.Province,
			PostalCode:  resp.Store.Address.PostalCode,
			Description: resp.Store.Address.Description},
		PhoneNumber: resp.Store.PhoneNumber,
		IsActive:    resp.Store.IsActive,
	}}, nil

}
