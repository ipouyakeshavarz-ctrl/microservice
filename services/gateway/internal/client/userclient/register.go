package userclient

import (
	"context"
	"myapp/api/gen/user"
	"myapp/pkg/richerror"

	"google.golang.org/grpc/status"
)

func (c *Client) Register(ctx context.Context, req *user.RegisterRequest) (*user.RegisterResponse, error) {
	const op = "userclient.Register"

	res, err := c.client.Register(ctx, req)
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
