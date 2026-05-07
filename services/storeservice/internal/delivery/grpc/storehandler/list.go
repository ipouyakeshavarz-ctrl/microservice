package storehandler

import (
	"context"
	"myapp/api/gen/store"

	"myapp/pkg/richerror"

	"storeapp/internal/param"
	"storeapp/internal/pkg/portobufermaper"
)

func (h *Handler) ListStoresByUser(ctx context.Context,
	req *store.ListStoresByUserRequest) (*store.ListStoresByUserResponse, error) {
	const op = "StoreHandler.ListStoresByUser"

	input := param.ListStoresByUserRequest{
		UserID: uint(req.UserId),
	}

	resp, err := h.storeSvc.ListStoresByUser(ctx, input)
	if err != nil {
		return nil, richerror.New(op).WithErr(err).WithFields(map[string]string{
			"massage": err.Error(),
		})
	}

	return portobufermaper.MapListResponseToProtobuf(resp), nil

}
