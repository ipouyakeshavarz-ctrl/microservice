package userclient

import (
	"context"
	"gatewayapp/internal/dto"
	"gatewayapp/internal/pkg/mapper"
	"myapp/pkg/richerror"

	"google.golang.org/grpc/status"
)

func (c *Client) Profile(ctx context.Context, req *dto.ProfileRequest) (*dto.ProfileResponse, error) {
	const op = "userclient.Profile"

	res, err := c.client.GetProfile(ctx, mapper.ToProfileRequest(req))
	if err != nil {
		if _, ok := status.FromError(err); ok {
			return nil, err
		}

		return nil, richerror.New(op).
			WithErr(err).
			WithKind(richerror.KindUnexpected)
	}

	return mapper.ToProfileResponse(res), nil
}
