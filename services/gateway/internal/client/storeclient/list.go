package storeclient

import (
	"context"
	"gatewayapp/internal/dto"
	"gatewayapp/internal/pkg/mapper"
	"myapp/pkg/richerror"

	"google.golang.org/grpc/status"
)

func (c *Client) ListStoresByUser(ctx context.Context, req *dto.ListStoresByUserRequest) (*dto.ListStoresByUserResponse, error) {
	const op = "storeclient.ListStoresByUser"

	res, err := c.client.ListStoresByUser(ctx, mapper.ToListStoresByUserRequest(req))
	if err != nil {
		if _, ok := status.FromError(err); ok {
			return nil, err
		}

		return nil, richerror.New(op).
			WithErr(err).
			WithKind(richerror.KindUnexpected)
	}
	return mapper.ToListStoresByUserResponse(res), nil
}
