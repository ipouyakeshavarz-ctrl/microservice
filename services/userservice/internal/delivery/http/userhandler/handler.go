package userhandler

import (
	authservice "userapp/internal/auth"
	"userapp/internal/service"
	"userapp/internal/validator"
)

type Handler struct {
	userSvc       userservice.Service
	userValidator validator.Validator
	authConfig    authservice.Config
	authSvc       authservice.Service
}

func New(userSvc userservice.Service,
	userValidator validator.Validator, authConfig authservice.Config,
	authSvc authservice.Service) Handler {
	return Handler{
		userSvc:       userSvc,
		userValidator: userValidator,
		authConfig:    authConfig,
		authSvc:       authSvc,
	}
}
