package userclient

import (
	"context"
	"myapp/api/gen/user"
	"myapp/pkg/richerror"

	"google.golang.org/grpc/status"
)

func (c *Client) Profile(ctx context.Context, req *user.ProfileRequest) (*user.ProfileResponse, error) {
	const op = "userclient.Profile"

	res, err := c.client.GetProfile(ctx, req)
	if err != nil {
		if _, ok := status.FromError(err); ok {
			return nil, err
		}

		return nil, richerror.New(op).
			WithErr(err).
			WithKind(richerror.KindUnexpected)
	}

	return res, nil
}
