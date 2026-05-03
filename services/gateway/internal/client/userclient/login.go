package userclient

import (
	"context"
	"myapp/api/gen/user"
	"myapp/pkg/richerror"
)

func (c *Client) Login(ctx context.Context, req *user.LoginRequest) (*user.LoginResponse, error) {
	const op = "userclient.Login"
	res, err := c.client.Login(ctx, req)
	if err != nil {
		return nil, richerror.New(op).WithErr(err)
	}

	return res, nil
}
