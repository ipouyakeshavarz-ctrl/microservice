package productclient

import (
	"context"
	"myapp/api/gen/product"
	"myapp/pkg/richerror"

	"google.golang.org/grpc/status"
)

func (c *Client) DeleteProduct(ctx context.Context, req *product.DeleteProductRequest) (*product.DeleteProductResponse, error) {
	const op = "productclient.DeleteProduct"

	res, err := c.client.DeleteProduct(ctx, req)
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
