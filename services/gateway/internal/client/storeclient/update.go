package storeclient

import (
	"context"
	"myapp/api/gen/store"
)

func (c *Client) UpdateStore(ctx context.Context,
	req *store.UpdateStoreRequest) (*store.UpdateStoreResponse, error) {

	res, err := c.client.UpdateStore(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
