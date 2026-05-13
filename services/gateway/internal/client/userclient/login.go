package userclient

import (
	"context"
	"gatewayapp/internal/dto"
	"gatewayapp/internal/pkg/mapper"
	"myapp/pkg/richerror"

	"google.golang.org/grpc/status"
)

func (c *Client) Login(ctx context.Context, req *dto.LoginRequest) (*dto.LoginResponse, error) {
	const op = "userclient.Login"
	res, err := c.client.Login(ctx, mapper.ToLoginRequest(req))
	if err != nil {
		if _, ok := status.FromError(err); ok {
			return nil, err
		}

		return nil, richerror.New(op).
			WithErr(err).
			WithKind(richerror.KindUnexpected)
	}

	return mapper.ToLoginResponse(res), nil
}
