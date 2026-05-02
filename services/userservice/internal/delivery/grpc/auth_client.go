package authclient

import (
	"context"
	authpb "myapp/api/gen/auth"
	"userapp/internal/domain"
)

type GRPCAuthClient struct {
	client authpb.AuthServiceClient
}

func NewGRPCAuthClient(client authpb.AuthServiceClient) *GRPCAuthClient {
	return &GRPCAuthClient{client: client}
}

func (a *GRPCAuthClient) GenerateTokens(ctx context.Context, user domain.User) (string, string, error) {

	resp, err := a.client.GenerateTokens(ctx, &authpb.UserInfo{
		Id:          uint64(user.ID),
		PhoneNumber: user.PhoneNumber,
		Name:        user.Name,
		Role:        domain.MapFromRoleEntity(user.Role),
	})

	if err != nil {
		return "", "", err
	}

	return resp.AccessToken, resp.RefreshToken, nil
}
