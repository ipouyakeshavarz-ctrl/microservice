package grpcserver

import (
	"authapp/internal/service"
	"context"
	"myapp/api/gen/auth"

	"google.golang.org/grpc"
)

type AuthGRPCServer struct {
	auth.UnimplementedAuthServiceServer
	authService *authservice.Service
}

func NewAuthGRPCServer(s *authservice.Service) *AuthGRPCServer {
	return &AuthGRPCServer{
		authService: s,
	}
}

func (g *AuthGRPCServer) GenerateTokens(ctx context.Context, req *auth.UserInfo, opts ...grpc.CallOption) (*auth.LoginTokenResponse, error) {

	resp, err := g.authService.GenerateTokens(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (g *AuthGRPCServer) ValidateTokens(ctx context.Context, req *auth.UserInfo, opts ...grpc.CallOption) (*auth.LoginTokenResponse, error) {
	resp, err := g.ValidateTokens(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
