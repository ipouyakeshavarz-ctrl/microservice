package producthandler

import (
	"myapp/api/gen/product"
	"productapp/internal/service"
)

type Handler struct {
	product.UnimplementedProductServiceServer
	productSvc productservice.Service
}

func New(productSvc productservice.Service) *Handler {
	return &Handler{
		productSvc: productSvc,
	}
}
