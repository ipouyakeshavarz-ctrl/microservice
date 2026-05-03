package productclient

import (
	"context"
	"myapp/api/gen/product"
)

func (c *Client) GetProductByID(ctx context.Context, req *product.GetProductByIDRequest) (*product.GetProductResponse, error) {

	res, err := c.client.GetProductByID(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
