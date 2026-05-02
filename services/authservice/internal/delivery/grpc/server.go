package grpcserver

import (
	"authapp/internal/service"
	"context"
	"myapp/api/gen/auth"
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

func (g *AuthGRPCServer) GenerateTokens(ctx context.Context, req *auth.UserInfo) (*auth.LoginTokenResponse, error) {

	resp, err := g.authService.GenerateTokens(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
