package productclient

import (
	"context"
	"myapp/api/gen/product"
)

func (c *Client) DeleteProduct(ctx context.Context, req *product.DeleteProductRequest) (*product.DeleteProductResponse, error) {

	res, err := c.client.DeleteProduct(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
