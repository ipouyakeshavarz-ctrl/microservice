package productclient

import (
	"context"
	"myapp/api/gen/product"
	"myapp/pkg/richerror"

	"google.golang.org/grpc/status"
)

func (c *Client) GetProductByID(ctx context.Context, req *product.GetProductByIDRequest) (*product.GetProductResponse, error) {
	const op = "productclient.GetProductByID"

	res, err := c.client.GetProductByID(ctx, req)
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
