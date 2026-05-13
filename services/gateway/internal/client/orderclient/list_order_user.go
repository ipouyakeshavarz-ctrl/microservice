package orderclient

import (
	"context"
	"gatewayapp/internal/dto"
	"gatewayapp/internal/pkg/mapper"
	"myapp/pkg/richerror"

	"google.golang.org/grpc/status"
)

func (c *Client) ListUserOrders(ctx context.Context, req *dto.ListUserOrdersRequest) (*dto.ListUserOrdersResponse, error) {
	const op = "orderClient.ListUserOrders"

	res, err := c.client.ListUserOrders(ctx, mapper.ToListUserOrdersRequest(req))
	if err != nil {
		if _, ok := status.FromError(err); ok {
			return nil, err
		}

		return nil, richerror.New(op).
			WithErr(err).
			WithKind(richerror.KindUnexpected)
	}
	return mapper.ToListUserOrdersResponse(res), nil
}
