package authservice

import (
	"authapp/internal/domain"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	jwt.RegisteredClaims
	UserID uint        `json:"user_id"`
	Role   domain.Role `json:"role"`
}

func (c Claims) Valid() error {
	return c.RegisteredClaims.Valid()
}
