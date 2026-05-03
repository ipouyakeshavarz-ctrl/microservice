package storeclient

import (
	"context"
	"myapp/api/gen/store"
	"myapp/pkg/richerror"
)

func (c *Client) UpdateStore(ctx context.Context,
	req *store.UpdateStoreRequest) (*store.UpdateStoreResponse, error) {
	const op = "storeclient.UpdateStore"

	res, err := c.client.UpdateStore(ctx, req)
	if err != nil {
		return nil, richerror.New(op).WithErr(err)
	}

	return res, nil
}
