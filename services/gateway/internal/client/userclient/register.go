package userclient

import (
	"context"
	"myapp/api/gen/user"
	"myapp/pkg/richerror"
)

func (c *Client) Register(ctx context.Context, req *user.RegisterRequest) (*user.RegisterResponse, error) {
	const op = "userclient.Register"

	res, err := c.client.Register(ctx, req)
	if err != nil {
		return nil, richerror.New(op).WithErr(err)
	}

	return res, nil
}
