package cartclient

import (
	"context"
	"myapp/api/gen/cart"
	"myapp/pkg/richerror"

	"google.golang.org/grpc/status"
)

func (c *Client) AddItem(ctx context.Context, req *cart.AddItemRequest) (*cart.AddItemResponse, error) {
	const op = "CartService.AddItem"
	res, err := c.client.AddItem(ctx, req)
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
