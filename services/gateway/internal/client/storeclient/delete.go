package storeclient

import (
	"context"
	"myapp/api/gen/store"
	"myapp/pkg/richerror"

	"google.golang.org/grpc/status"
)

func (c *Client) DeleteStore(ctx context.Context, req *store.DeleteStoreRequest) (*store.DeleteStoreResponse, error) {
	const op = "storeclient.DeleteStore"

	res, err := c.client.DeleteStore(ctx, req)
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
