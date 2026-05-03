package userhandler

import (
	"myapp/api/gen/user"
	"userapp/internal/service"
	"userapp/internal/validator"
)

type Handler struct {
	user.UnimplementedUserServiceServer
	userSvc       userservice.Service
	userValidator validator.Validator
}

func New(userSvc userservice.Service,
	userValidator validator.Validator) *Handler {
	return &Handler{
		userSvc:       userSvc,
		userValidator: userValidator,
	}
}
