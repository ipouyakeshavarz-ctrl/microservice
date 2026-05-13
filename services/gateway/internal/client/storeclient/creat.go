package storeclient

import (
	"context"
	"gatewayapp/internal/dto"
	"gatewayapp/internal/pkg/mapper"
	"myapp/pkg/richerror"

	"google.golang.org/grpc/status"
)

func (c *Client) CreateStore(ctx context.Context,
	req *dto.CreateStoreRequest) (*dto.CreateStoreResponse, error) {
	const op = "storeclient.CreateStore"

	res, err := c.client.CreateStore(ctx,mapper.ToCreateStoreRequest(req))
	if err != nil {
		if _, ok := status.FromError(err); ok {
			return nil, err
		}

		return nil, richerror.New(op).
			WithErr(err).
			WithKind(richerror.KindUnexpected)
	}

	return mapper.ToCreateStoreResponse(res), nil
}
