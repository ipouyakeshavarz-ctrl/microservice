package userclient

import (
	"context"
	"myapp/api/gen/user"
)

func (c *Client) Register(ctx context.Context, req *user.RegisterRequest) (*user.RegisterResponse, error) {

	res, err := c.client.Register(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
