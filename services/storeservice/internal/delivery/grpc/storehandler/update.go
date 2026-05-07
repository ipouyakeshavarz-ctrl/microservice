package storehandler

import (
	"context"
	"myapp/api/gen/store"
	"myapp/pkg/errmsg"
	"myapp/pkg/richerror"
	"storeapp/internal/domain"
	"storeapp/internal/param"
)

func (h *Handler) UpdateStore(ctx context.Context, req *store.UpdateStoreRequest) (*store.UpdateStoreResponse, error) {
	const op = "StoreHandler.UpdateStore"

	input := param.UpdateStoreRequest{
		ID:          uint(req.StoreId),
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
		IsActive:    req.IsActive,
	}

	if fieldErrors, err := h.storeValidator.ValidateUpdateRequest(input); err != nil {
		return nil, richerror.New(op).
			WithKind(richerror.KindInvalid).
			WithMessage(errmsg.ErrorMsgInvalidInput).
			WithFields(fieldErrors)
	}

	resp, err := h.storeSvc.UpdateStore(ctx, input)
	if err != nil {
		return nil, richerror.New(op).WithErr(err).WithFields(map[string]string{
			"massage": err.Error(),
		})
	}

	return &store.UpdateStoreResponse{
		Store: &store.StoreInfo{
			Id:          uint64(resp.Store.UserID),
			UserId:      uint64(resp.Store.UserID),
			Name:        resp.Store.Name,
			Description: resp.Store.Description,
			LogoUrl:     resp.Store.LogoURL,
			Address: &store.Address{
				Street:      resp.Store.Address.Street,
				City:        resp.Store.Address.City,
				Province:    resp.Store.Address.Province,
				PostalCode:  resp.Store.Address.PostalCode,
				Description: resp.Store.Address.Description,
			},
			PhoneNumber: resp.Store.PhoneNumber,
			IsActive:    resp.Store.IsActive,
		}}, nil
}
