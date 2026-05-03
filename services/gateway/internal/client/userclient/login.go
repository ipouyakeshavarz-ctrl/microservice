package userclient

import (
	"context"
	"myapp/api/gen/user"
)

func (c *Client) Login(ctx context.Context, req *user.LoginRequest) (*user.LoginResponse, error) {

	res, err := c.client.Login(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
