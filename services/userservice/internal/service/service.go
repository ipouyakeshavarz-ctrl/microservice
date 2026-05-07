package userservice

import (
	"context"

	"userapp/internal/domain"

	"golang.org/x/crypto/bcrypt"
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

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
