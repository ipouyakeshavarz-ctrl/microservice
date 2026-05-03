package storeclient

import (
	"context"
	"myapp/api/gen/store"
)

func (c *Client) ListStoresByUser(ctx context.Context, req *store.ListStoresByUserRequest) (*store.ListStoresByUserResponse, error) {

	res, err := c.client.ListStoresByUser(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
