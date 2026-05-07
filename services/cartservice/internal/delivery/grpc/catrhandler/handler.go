package carthandler

import (
	cartservice "cartapp/internal/service"
	"myapp/api/gen/cart"
)

type Handler struct {
	cart.UnimplementedCartServiceServer
	cartSvc cartservice.Service
}

func New(cartSvc cartservice.Service) *Handler {
	return &Handler{
		cartSvc: cartSvc,
	}
}
