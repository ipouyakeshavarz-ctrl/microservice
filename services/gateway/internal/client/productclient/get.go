package productclient

import (
	"context"
	"myapp/api/gen/product"
	"myapp/pkg/richerror"
)

func (c *Client) GetProductByID(ctx context.Context, req *product.GetProductByIDRequest) (*product.GetProductResponse, error) {
	const op = "productclient.GetProductByID"

	res, err := c.client.GetProductByID(ctx, req)
	if err != nil {
		return nil, richerror.New(op).WithErr(err)
	}

	return res, nil
}
