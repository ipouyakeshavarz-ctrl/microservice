package productclient

import (
	"myapp/api/gen/product"
	"myapp/pkg/richerror"

	"google.golang.org/grpc"
)

type Client struct {
	conn   *grpc.ClientConn
	client product.ProductServiceClient
}

func New(addr string) (*Client, error) {
	const op = "productclient.New"
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, richerror.New(op).WithErr(err)
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
