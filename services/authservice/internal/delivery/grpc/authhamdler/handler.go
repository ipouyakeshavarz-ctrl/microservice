package authhandler

import (
	authservice "authapp/internal/service"
	"myapp/api/gen/auth"
)

type Handler struct {
	auth.UnimplementedAuthServiceServer
	authService *authservice.Service
}

func New(s *authservice.Service) *Handler {
	return &Handler{
		authService: s,
	}
}
