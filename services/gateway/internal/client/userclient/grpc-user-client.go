package userclient

import (
	"myapp/api/gen/user"

	"google.golang.org/grpc"
)

type Client struct {
	conn   *grpc.ClientConn
	client user.UserServiceClient
}

func New(addr string) (*Client, error) {

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	c := user.NewUserServiceClient(conn)

	return &Client{
		conn:   conn,
		client: c,
	}, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}
