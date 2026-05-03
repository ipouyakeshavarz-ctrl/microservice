package storehandler

import (
	"myapp/api/gen/store"
	"storeapp/internal/service"
	"storeapp/internal/validator"
)

type Handler struct {
	store.UnimplementedStoreServiceServer
	storeSvc       storeservice.Service
	storeValidator storevalidator.Validator
}

func New(storeSvc storeservice.Service, storeValidator storevalidator.Validator) *Handler {
	return &Handler{
		storeSvc:       storeSvc,
		storeValidator: storeValidator,
	}
}
