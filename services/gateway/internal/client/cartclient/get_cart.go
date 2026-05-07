package cartclient

import (
	"context"
	"myapp/api/gen/cart"
	"myapp/pkg/richerror"

	"google.golang.org/grpc/status"
)

func (c *Client) GetCart(ctx context.Context, req *cart.GetCartRequest) (*cart.GetCartResponse, error) {
	const op = "CartService.GetCart"
	res, err := c.client.GetCart(ctx, req)
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
