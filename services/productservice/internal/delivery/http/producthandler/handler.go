package producthandler

import (
	"productapp/internal/auth"
	"productapp/internal/service"
)

type Handler struct {
	authConfig authservice.Config
	authSvc    authservice.Service
	productSvc productservice.Service
}

func New(authConfig authservice.Config, authSvc authservice.Service, productSvc productservice.Service) Handler {
	return Handler{
		authConfig: authConfig,
		authSvc:    authSvc,
		productSvc: productSvc,
	}
}
