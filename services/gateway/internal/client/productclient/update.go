package productclient

import (
	"context"
	"myapp/api/gen/product"
)

func (c *Client) UpdateProduct(ctx context.Context,
	req *product.UpdateProductRequest) (*product.UpdateProductResponse, error) {

	res, err := c.client.UpdateProduct(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
