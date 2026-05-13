package productclient

import (
	"context"
	"gatewayapp/internal/dto"
	"gatewayapp/internal/pkg/mapper"
	"myapp/pkg/richerror"

	"google.golang.org/grpc/status"
)

func (c *Client) CreateProduct(ctx context.Context,
	req *dto.CreateProductRequest) (*dto.CreateProductResponse, error) {
	const op = "productclient.CreateProduct"
	res, err := c.client.CreateProduct(ctx, mapper.ToCreateProductRequest(req))
	if err != nil {
		if _, ok := status.FromError(err); ok {
			return nil, err
		}

		return nil, richerror.New(op).
			WithErr(err).
			WithKind(richerror.KindUnexpected)
	}
	return mapper.ToCreateProductResponse(res), nil
}
