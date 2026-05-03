package productclient

import (
	"context"
	"myapp/api/gen/product"
	"myapp/pkg/richerror"
)

func (c *Client) DeleteProduct(ctx context.Context, req *product.DeleteProductRequest) (*product.DeleteProductResponse, error) {
	const op = "productclient.DeleteProduct"

	res, err := c.client.DeleteProduct(ctx, req)
	if err != nil {
		return nil, richerror.New(op).WithErr(err)
	}

	return res, nil
}
