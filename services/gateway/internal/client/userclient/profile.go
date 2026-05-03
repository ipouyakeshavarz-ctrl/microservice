package userclient

import (
	"context"
	"myapp/api/gen/user"
)

func (c *Client) Profile(ctx context.Context, req *user.ProfileRequest) (*user.ProfileResponse, error) {

	res, err := c.client.GetProfile(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
