package storeclient

import (
	"context"
	"myapp/api/gen/store"
	"myapp/pkg/richerror"
)

func (c *Client) CreateStore(ctx context.Context,
	req *store.CreateStoreRequest) (*store.CreateStoreResponse, error) {
	const op = "storeclient.CreateStore"

	res, err := c.client.CreateStore(ctx, req)
	if err != nil {
		return nil, richerror.New(op).WithErr(err)
	}

	return res, nil
}
