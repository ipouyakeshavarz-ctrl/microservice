package userclient

import (
	"context"
	"myapp/api/gen/user"
	"myapp/pkg/richerror"

	"google.golang.org/grpc/status"
)

func (c *Client) Login(ctx context.Context, req *user.LoginRequest) (*user.LoginResponse, error) {
	const op = "userclient.Login"
	res, err := c.client.Login(ctx, req)
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
