package storeclient

import (
	"context"
	"myapp/api/gen/store"
	"myapp/pkg/richerror"

	"google.golang.org/grpc/status"
)

func (c *Client) CreateStore(ctx context.Context,
	req *store.CreateStoreRequest) (*store.CreateStoreResponse, error) {
	const op = "storeclient.CreateStore"

	res, err := c.client.CreateStore(ctx, req)
	if err != nil {
		if _, ok := status.FromError(err); ok {
			return nil, err
		}

		return nil, richerror.New(op).
			WithErr(err).
			WithKind(richerror.KindUnexpected)
	}

	return res, nil
}
