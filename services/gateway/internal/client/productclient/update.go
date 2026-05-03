package productclient

import (
	"context"
	"myapp/api/gen/product"
	"myapp/pkg/richerror"
)

func (c *Client) UpdateProduct(ctx context.Context,
	req *product.UpdateProductRequest) (*product.UpdateProductResponse, error) {
	const op = "productclient.UpdateProduct"

	res, err := c.client.UpdateProduct(ctx, req)
	if err != nil {
		return nil, richerror.New(op).WithErr(err)
	}

	return res, nil
}
