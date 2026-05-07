package storeclient

import (
	"context"
	"myapp/api/gen/store"
	"myapp/pkg/richerror"

	"google.golang.org/grpc/status"
)

func (c *Client) UpdateStore(ctx context.Context,
	req *store.UpdateStoreRequest) (*store.UpdateStoreResponse, error) {
	const op = "storeclient.UpdateStore"

	res, err := c.client.UpdateStore(ctx, req)
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
