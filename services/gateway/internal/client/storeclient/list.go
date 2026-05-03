package storeclient

import (
	"context"
	"myapp/api/gen/store"
	"myapp/pkg/richerror"
)

func (c *Client) ListStoresByUser(ctx context.Context, req *store.ListStoresByUserRequest) (*store.ListStoresByUserResponse, error) {
	const op = "storeclient.ListStoresByUser"

	res, err := c.client.ListStoresByUser(ctx, req)
	if err != nil {
		return nil, richerror.New(op).WithErr(err)
	}

	return res, nil
}
