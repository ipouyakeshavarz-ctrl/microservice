package orderclient

import (
	"myapp/api/gen/order"
	"myapp/pkg/richerror"

	"google.golang.org/grpc"
)

type Client struct {
	conn   *grpc.ClientConn
	client order.OrderServiceClient
}

func New(addr string) (*Client, error) {
	const op = "orderclient.New"
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, richerror.New(op).WithErr(err)
	}

	c := order.NewOrderServiceClient(conn)

	return &Client{
		conn:   conn,
		client: c,
	}, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}
