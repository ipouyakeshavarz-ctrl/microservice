package authclient

import (
	"context"
	"myapp/api/gen/auth"
	"myapp/pkg/richerror"

	"google.golang.org/grpc"
)

type Client struct {
	conn   *grpc.ClientConn
	client auth.AuthServiceClient
}

func New(addr string) (*Client, error) {
	const op = "authclient.New"

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, richerror.New(op).WithErr(err)
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
	const op = "authclient.GenerateTokens"
	res, err := c.client.GenerateTokens(ctx, req)
	if err != nil {
		return nil, richerror.New(op).WithErr(err)
	}

	return res, nil
}
func (c *Client) VerifyToken(ctx context.Context, req *auth.VerifyTokenRequest, opts ...grpc.CallOption) (*auth.VerifyTokenResponse, error) {
	const op = "authclient.VerifyToken"
	res, err := c.client.VerifyToken(ctx, req)
	if err != nil {
		return nil, richerror.New(op).WithErr(err)
	}

	return res, nil
}
