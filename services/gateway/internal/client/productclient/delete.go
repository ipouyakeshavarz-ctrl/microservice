package productclient

import (
	"context"
	"gatewayapp/internal/dto"
	"gatewayapp/internal/pkg/mapper"
	"myapp/pkg/richerror"

	"google.golang.org/grpc/status"
)

func (c *Client) DeleteProduct(ctx context.Context, req *dto.DeleteProductRequest) (*dto.DeleteProductResponse, error) {
	const op = "productclient.DeleteProduct"

	res, err := c.client.DeleteProduct(ctx, mapper.ToDeleteProductRequest(req))
	if err != nil {
		if _, ok := status.FromError(err); ok {
			return nil, err
		}

		return nil, richerror.New(op).
			WithErr(err).
			WithKind(richerror.KindUnexpected)
	}
	return mapper.ToDeleteProductResponse(res), nil
}
