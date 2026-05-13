package cartclient

import (
	"context"
	"gatewayapp/internal/dto"
	"gatewayapp/internal/pkg/mapper"
	"myapp/pkg/richerror"

	"google.golang.org/grpc/status"
)

func (c *Client) AddItem(ctx context.Context, req *dto.AddItemRequest) (*dto.AddItemResponse, error) {
	const op = "CartService.AddItem"
	res, err := c.client.AddItem(ctx, mapper.ToAddItemRequest(req))
	if err != nil {
		if _, ok := status.FromError(err); ok {
			return nil, err
		}

		return nil, richerror.New(op).
			WithErr(err).
			WithKind(richerror.KindUnexpected)
	}
	return mapper.ToAddItemResponse(res), nil

}
