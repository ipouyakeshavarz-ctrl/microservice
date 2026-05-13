package cartclient

import (
	"context"
	"gatewayapp/internal/dto"
	"gatewayapp/internal/pkg/mapper"
	"myapp/pkg/richerror"

	"google.golang.org/grpc/status"
)

func (c *Client) GetCart(ctx context.Context, req *dto.GetCartRequest) (*dto.GetCartResponse, error) {
	const op = "CartService.GetCart"
	res, err := c.client.GetCart(ctx, mapper.ToGetCartRequest(req))
	if err != nil {
		if _, ok := status.FromError(err); ok {
			return nil, err
		}

		return nil, richerror.New(op).
			WithErr(err).
			WithKind(richerror.KindUnexpected)
	}

	return mapper.ToGetCartResponse(res), nil
}
