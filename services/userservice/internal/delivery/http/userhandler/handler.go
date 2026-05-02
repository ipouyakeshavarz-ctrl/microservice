package userhandler

import (
	"userapp/internal/service"
	"userapp/internal/validator"
)

type Handler struct {
	userSvc       userservice.Service
	userValidator validator.Validator
}

func New(userSvc userservice.Service,
	userValidator validator.Validator) Handler {
	return Handler{
		userSvc:       userSvc,
		userValidator: userValidator,
	}
}
