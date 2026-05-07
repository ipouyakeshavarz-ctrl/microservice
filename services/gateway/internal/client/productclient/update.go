package productclient

import (
	"context"
	"myapp/api/gen/product"
	"myapp/pkg/richerror"

	"google.golang.org/grpc/status"
)

func (c *Client) UpdateProduct(ctx context.Context,
	req *product.UpdateProductRequest) (*product.UpdateProductResponse, error) {
	const op = "productclient.UpdateProduct"

	res, err := c.client.UpdateProduct(ctx, req)
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
