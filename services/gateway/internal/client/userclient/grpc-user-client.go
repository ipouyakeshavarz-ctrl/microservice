package userclient

import (
	"myapp/api/gen/user"
	"myapp/pkg/richerror"

	"google.golang.org/grpc"
)

type Client struct {
	conn   *grpc.ClientConn
	client user.UserServiceClient
}

func New(addr string) (*Client, error) {
	const op = "userclient.New"

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, richerror.New(op).WithErr(err)
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
