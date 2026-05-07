package storehandler

import (
	"context"
	"myapp/api/gen/store"
	"myapp/pkg/errmsg"
	"myapp/pkg/richerror"
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
		if re, ok := err.(*richerror.RichError); ok {
			return nil, re
		}

		return nil, richerror.New(op).
			WithKind(richerror.KindUnexpected).
			WithMessage(errmsg.ErrorMsgFailedToDeleteStore).
			WithErr(err)
	}

	return &store.DeleteStoreResponse{
		Success: resp.Success,
	}, nil

}
