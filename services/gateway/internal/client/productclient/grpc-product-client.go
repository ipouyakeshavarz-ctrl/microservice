package productclient

import (
	"myapp/api/gen/product"

	"google.golang.org/grpc"
)

type Client struct {
	conn   *grpc.ClientConn
	client product.ProductServiceClient
}

func New(addr string) (*Client, error) {

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	c := product.NewProductServiceClient(conn)

	return &Client{
		conn:   conn,
		client: c,
	}, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}
