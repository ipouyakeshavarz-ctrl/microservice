package userservice

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"userapp/internal/domain"
)

type Repository interface {
	Register(ctx context.Context, u domain.User) (domain.User, error)
	GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (domain.User, error)
	GetUserByID(ctx context.Context, userID uint) (domain.User, error)
}

type AuthClient interface {
	GenerateTokens(ctx context.Context, user domain.User) (accessToken string, refreshToken string, err error)
}

type Service struct {
	authClient AuthClient
	repo       Repository
}

func New(authClient AuthClient, repo Repository) Service {
	return Service{authClient: authClient, repo: repo}
}

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
