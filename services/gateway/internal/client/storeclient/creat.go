package storeclient

import (
	"context"
	"myapp/api/gen/store"
)

func (c *Client) CreateStore(ctx context.Context,
	req *store.CreateStoreRequest) (*store.CreateStoreResponse, error) {

	res, err := c.client.CreateStore(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
