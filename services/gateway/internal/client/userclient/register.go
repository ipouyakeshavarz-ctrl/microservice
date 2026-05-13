package userclient

import (
	"context"
	"gatewayapp/internal/dto"
	"gatewayapp/internal/pkg/mapper"
	"myapp/pkg/richerror"

	"google.golang.org/grpc/status"
)

func (c *Client) Register(ctx context.Context, req *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	const op = "userclient.Register"

	res, err := c.client.Register(ctx, mapper.ToRegisterRequest(req))
	if err != nil {
		if _, ok := status.FromError(err); ok {
			return nil, err
		}

		return nil, richerror.New(op).
			WithErr(err).
			WithKind(richerror.KindUnexpected)
	}

	return mapper.ToRegisterResponse(res), nil
}
