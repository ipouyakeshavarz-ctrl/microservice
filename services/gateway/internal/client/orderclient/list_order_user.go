package orderclient

import (
	"context"
	"myapp/api/gen/order"
	"myapp/pkg/richerror"

	"google.golang.org/grpc/status"
)

func (c *Client) ListUserOrders(ctx context.Context,
	req *order.ListUserOrdersRequest) (*order.ListUserOrdersResponse, error) {
	const op = "orderClient.ListUserOrders"

	res, err := c.client.ListUserOrders(ctx, req)
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
