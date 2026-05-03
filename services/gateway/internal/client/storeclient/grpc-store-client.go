package storeclient

import (
	"myapp/api/gen/store"

	"google.golang.org/grpc"
)

type Client struct {
	conn   *grpc.ClientConn
	client store.StoreServiceClient
}

func New(addr string) (*Client, error) {

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	c := store.NewStoreServiceClient(conn)

	return &Client{
		conn:   conn,
		client: c,
	}, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}
