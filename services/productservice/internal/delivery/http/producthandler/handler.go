package producthandler

import (
	"productapp/internal/service"
)

type Handler struct {
	productSvc productservice.Service
}

func New(productSvc productservice.Service) Handler {
	return Handler{
		productSvc: productSvc,
	}
}
