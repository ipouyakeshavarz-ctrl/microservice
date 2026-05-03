package userclient

import (
	"context"
	"myapp/api/gen/user"
	"myapp/pkg/richerror"
)

func (c *Client) Profile(ctx context.Context, req *user.ProfileRequest) (*user.ProfileResponse, error) {
	const op = "userclient.Profile"

	res, err := c.client.GetProfile(ctx, req)
	if err != nil {
		return nil, richerror.New(op).WithErr(err)
	}

	return res, nil
}
