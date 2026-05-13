package storeclient

import (
	"context"
	"gatewayapp/internal/dto"
	"gatewayapp/internal/pkg/mapper"
	"myapp/pkg/richerror"

	"google.golang.org/grpc/status"
)

func (c *Client) DeleteStore(ctx context.Context, req *dto.DeleteStoreRequest) (*dto.DeleteStoreResponse, error) {
	const op = "storeclient.DeleteStore"

	res, err := c.client.DeleteStore(ctx, mapper.ToDeleteStoreRequest(req))
	if err != nil {
		if _, ok := status.FromError(err); ok {
			return nil, err
		}

		return nil, richerror.New(op).
			WithErr(err).
			WithKind(richerror.KindUnexpected)
	}

	return mapper.ToDeleteStoreResponse(res), nil
}
