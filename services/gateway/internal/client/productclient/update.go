package productclient

import (
	"context"
	"gatewayapp/internal/dto"
	"gatewayapp/internal/pkg/mapper"
	"myapp/pkg/richerror"

	"google.golang.org/grpc/status"
)

func (c *Client) UpdateProduct(ctx context.Context,
	req *dto.UpdateProductRequest) (*dto.UpdateProductResponse, error) {
	const op = "productclient.UpdateProduct"

	res, err := c.client.UpdateProduct(ctx, mapper.ToUpdateProductRequest(req))
	if err != nil {
		if _, ok := status.FromError(err); ok {
			return nil, err
		}

		return nil, richerror.New(op).
			WithErr(err).
			WithKind(richerror.KindUnexpected)
	}
	return mapper.ToUpdateProductResponse(res), nil
}
