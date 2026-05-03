package storeclient

import (
	"context"
	"myapp/api/gen/store"
)

func (c *Client) DeleteStore(ctx context.Context, req *store.DeleteStoreRequest) (*store.DeleteStoreResponse, error) {

	res, err := c.client.DeleteStore(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
