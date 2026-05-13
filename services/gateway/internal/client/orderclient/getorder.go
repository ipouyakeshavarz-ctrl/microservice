package orderclient

import (
	"context"
	"gatewayapp/internal/dto"
	"gatewayapp/internal/pkg/mapper"
	"myapp/pkg/richerror"

	"google.golang.org/grpc/status"
)

func (c *Client) GetOrder(ctx context.Context, req *dto.GetOrderRequest) (*dto.GetOrderResponse, error) {
	const op = "orderClient.GetOrder"

	res, err := c.client.GetOrder(ctx, mapper.ToGetOrderRequest(req))
	if err != nil {
		if _, ok := status.FromError(err); ok {
			return nil, err
		}

		return nil, richerror.New(op).
			WithErr(err).
			WithKind(richerror.KindUnexpected)
	}
	return mapper.ToGetOrderResponse(res), nil
}
