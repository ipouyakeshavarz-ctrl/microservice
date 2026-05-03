package productclient

import (
	"context"
	"myapp/api/gen/product"
	"myapp/pkg/richerror"
)

func (c *Client) CreateProduct(ctx context.Context,
	req *product.CreateProductRequest) (*product.CreateProductResponse, error) {
	const op = "productclient.CreateProduct"

	res, err := c.client.CreateProduct(ctx, req)
	if err != nil {
		return nil, richerror.New(op).WithErr(err)
	}

	return res, nil
}
