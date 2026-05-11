package orderhandler

import (
	"myapp/api/gen/order"
	orderservice "orderapp/internal/service"
)

type Handler struct {
	order.UnimplementedOrderServiceServer
	orderSvc orderservice.Service
}

func New(orderSvc orderservice.Service) *Handler {
	return &Handler{
		orderSvc: orderSvc,
	}
}
