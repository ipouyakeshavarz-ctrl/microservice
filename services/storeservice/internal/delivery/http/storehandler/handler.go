package storehandler

import (
	"storeapp/internal/auth"
	"storeapp/internal/service"
	"storeapp/internal/validator"
)

type Handler struct {
	authConfig     authservice.Config
	authSvc        authservice.Service
	storeSvc       storeservice.Service
	storeValidator storevalidator.Validator
}

func New(authConfig authservice.Config, authSvc authservice.Service, storeSvc storeservice.Service, storeValidator storevalidator.Validator) Handler {
	return Handler{
		authConfig:     authConfig,
		authSvc:        authSvc,
		storeSvc:       storeSvc,
		storeValidator: storeValidator,
	}
}
