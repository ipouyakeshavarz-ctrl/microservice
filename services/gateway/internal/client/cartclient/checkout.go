package cartclient

import (
	"context"
	"gatewayapp/internal/dto"
	"gatewayapp/internal/pkg/mapper"
	"myapp/pkg/richerror"

	"google.golang.org/grpc/status"
)

func (c *Client) Checkout(ctx context.Context, req *dto.CheckoutRequest) (*dto.CheckoutResponse, error) {
	const op = "cartservice.Checkout"

	res, err := c.client.Checkout(ctx, mapper.ToCheckoutRequest(req))
	if err != nil {
		if _, ok := status.FromError(err); ok {
			return nil, err
		}

		return nil, richerror.New(op).
			WithErr(err).
			WithKind(richerror.KindUnexpected)
	}
	return mapper.ToCheckoutResponse(res), nil

}
