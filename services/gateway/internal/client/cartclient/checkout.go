package cartclient

import (
	"context"
	"myapp/api/gen/cart"
	"myapp/pkg/richerror"

	"google.golang.org/grpc/status"
)

func (c *Client) Checkout(ctx context.Context,
	req *cart.CheckoutRequest) (*cart.CheckoutResponse, error) {
	const op = "cartservice.Checkout"

	res, err := c.client.Checkout(ctx, req)
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
