package storeclient

import (
	"context"
	"myapp/api/gen/store"
	"myapp/pkg/richerror"
)

func (c *Client) DeleteStore(ctx context.Context, req *store.DeleteStoreRequest) (*store.DeleteStoreResponse, error) {
	const op = "storeclient.DeleteStore"

	res, err := c.client.DeleteStore(ctx, req)
	if err != nil {
		return nil, richerror.New(op).WithErr(err)
	}

	return res, nil
}
