package productclient

import (
	"context"
	"myapp/api/gen/product"
	"myapp/pkg/richerror"

	"google.golang.org/grpc/status"
)

func (c *Client) CreateProduct(ctx context.Context,
	req *product.CreateProductRequest) (*product.CreateProductResponse, error) {
	const op = "productclient.CreateProduct"
	res, err := c.client.CreateProduct(ctx, req)
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
