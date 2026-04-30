package userservice

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"userapp/internal/entity"
)

type Repository interface {
	Register(ctx context.Context, u entity.User) (entity.User, error)
	GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (entity.User, error)
	GetUserByID(ctx context.Context, userID uint) (entity.User, error)
}

// AuthGenerator todo if need add ctx to this interface
type AuthGenerator interface {
	CreateAccessToken(user entity.User) (string, error)
	CreateRefreshToken(user entity.User) (string, error)
}

type Service struct {
	auth AuthGenerator
	repo Repository
}

func New(auth AuthGenerator, repo Repository) Service {
	return Service{auth: auth, repo: repo}
}

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
