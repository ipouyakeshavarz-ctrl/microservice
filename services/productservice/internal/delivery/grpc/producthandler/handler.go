package producthandler

import (
	"myapp/api/gen/product"
	"productapp/internal/service"
	productvalidator "productapp/internal/validator"
)

type Handler struct {
	product.UnimplementedProductServiceServer
	productSvc productservice.Service
	productV   productvalidator.Validator
}

func New(productSvc productservice.Service, productV productvalidator.Validator) *Handler {
	return &Handler{
		productSvc: productSvc,
		productV:   productV,
	}
}
