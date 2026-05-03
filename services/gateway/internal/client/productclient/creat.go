package productclient

import (
	"context"
	"myapp/api/gen/product"
)

func (c *Client) CreateProduct(ctx context.Context,
	req *product.CreateProductRequest) (*product.CreateProductResponse, error) {

	res, err := c.client.CreateProduct(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
