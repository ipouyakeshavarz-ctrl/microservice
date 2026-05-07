package storeclient

import (
	"context"
	"myapp/api/gen/store"
	"myapp/pkg/richerror"

	"google.golang.org/grpc/status"
)

func (c *Client) ListStoresByUser(ctx context.Context, req *store.ListStoresByUserRequest) (*store.ListStoresByUserResponse, error) {
	const op = "storeclient.ListStoresByUser"

	res, err := c.client.ListStoresByUser(ctx, req)
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
