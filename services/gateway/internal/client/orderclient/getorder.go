package orderclient

import (
	"context"
	"myapp/api/gen/order"
	"myapp/pkg/richerror"

	"google.golang.org/grpc/status"
)

func (c *Client) GetOrder(ctx context.Context,
	req *order.GetOrderRequest) (*order.GetOrderResponse, error) {
	const op = "orderClient.GetOrder"

	res, err := c.client.GetOrder(ctx, req)
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
