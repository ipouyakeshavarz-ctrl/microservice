package authgrpc

import "google.golang.org/grpc"

type AuthClient struct {
	conn *grpc.ClientConn
}

func NewAuthClient(conn *grpc.ClientConn) *AuthClient {
	return &AuthClient{conn: conn}
}

func (a *AuthClient) Login(email, password string) (string, error) {

	// TODO: call grpc auth service

	return "dummy-token", nil
}

func (a *AuthClient) ValidateToken(token string) (string, error) {

	// TODO: grpc validate

	return "1", nil
}
