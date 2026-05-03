package storehandler

import (
	"context"
	"myapp/api/gen/store"
	"storeapp/internal/param"
)

func (h *Handler) DeleteStore(ctx context.Context,
	req *store.DeleteStoreRequest) (*store.DeleteStoreResponse, error) {
	const op = "storeHandler.DeleteStore"

	input := param.DeleteStoreRequest{
		StoreID: uint(req.StoreId),
		UserID:  uint(req.UserId),
	}

	resp, err := h.storeSvc.DeleteStore(ctx, input)
	if err != nil {
		return nil, err
	}

	return &store.DeleteStoreResponse{
		Success: resp.Success,
	}, nil

}
