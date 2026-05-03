package authclient

import (
	"context"
	"myapp/api/gen/auth"

	"google.golang.org/grpc"
)

type Client struct {
	conn   *grpc.ClientConn
	client auth.AuthServiceClient
}

func New(addr string) (*Client, error) {

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	c := auth.NewAuthServiceClient(conn)

	return &Client{
		conn:   conn,
		client: c,
	}, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}

func (c *Client) GenerateTokens(ctx context.Context, req *auth.UserInfo, opts ...grpc.CallOption) (*auth.LoginTokenResponse, error) {
	res, err := c.client.GenerateTokens(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
func (c *Client) VerifyToken(ctx context.Context, req *auth.VerifyTokenRequest, opts ...grpc.CallOption) (*auth.VerifyTokenResponse, error) {
	res, err := c.client.VerifyToken(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
