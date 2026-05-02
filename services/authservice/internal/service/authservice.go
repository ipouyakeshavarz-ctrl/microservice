package authservice

import (
	"authapp/internal/domain"
	"context"
	"myapp/api/gen/auth"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Config struct {
	SignKey               string        `koanf:"sign_key"`
	AccessExpirationTime  time.Duration `koanf:"access_expiration_time"`
	RefreshExpirationTime time.Duration `koanf:"refresh_expiration_time"`
	AccessSubject         string        `koanf:"access_subject"`
	RefreshSubject        string        `koanf:"refresh_subject"`
}

type Service struct {
	auth.UnimplementedAuthServiceServer
	config Config
}

func New(cfg Config) *Service {
	return &Service{
		config: cfg,
	}
}

func (s *Service) GenerateTokens(ctx context.Context, u *auth.UserInfo) (*auth.LoginTokenResponse, error) {

	user := domain.User{
		ID:          uint(u.Id),
		PhoneNumber: u.PhoneNumber,
		Name:        u.Name,
		Role:        domain.MapToRoleEntity(u.Role),
	}

	accessToken, err := s.CreateAccessToken(user)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.CreateRefreshToken(user)
	if err != nil {
		return nil, err
	}

	// save refresh token for logout / refresh support
	//err = s.sessionRepo.SaveRefreshToken(ctx, user.ID, refreshToken, a.jwt.Config.RefreshExpirationTime)
	//if err != nil {
	//	return nil, err
	//}

	return &auth.LoginTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    int64(s.config.AccessExpirationTime.Seconds()),
	}, nil
}

func (s Service) CreateAccessToken(user domain.User) (string, error) {
	return s.createToken(user.ID, user.Role, s.config.AccessSubject, s.config.AccessExpirationTime)
}

func (s Service) CreateRefreshToken(user domain.User) (string, error) {
	return s.createToken(user.ID, user.Role, s.config.RefreshSubject, s.config.RefreshExpirationTime)
}

func (s Service) ParseToken(bearerToken string) (*Claims, error) {

	tokenStr := strings.Replace(bearerToken, "Bearer ", "", 1)

	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.config.SignKey), nil

	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}

}

func (s Service) createToken(userID uint, role domain.Role, subject string, expireDuration time.Duration) (string, error) {
	// create a signer for rsa 256
	// TODO - replace with rsa 256 RS256 - https://github.com/golang-jwt/jwt/blob/main/http_example_test.go

	// set our claims
	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   subject,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireDuration)),
		},
		UserID: userID,
		Role:   role,
	}

	// TODO - add sign method to config
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := accessToken.SignedString([]byte(s.config.SignKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
