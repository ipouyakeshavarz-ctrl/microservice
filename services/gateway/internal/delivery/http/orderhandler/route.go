package orderhandler

import (
	"gatewayapp/internal/delivery/http/middleware"

	"github.com/labstack/echo/v4"
)

func (h Handler) SetRoutes(e *echo.Echo) {
	storeGroup := e.Group("/order")

	storeGroup.GET("/:id", h.GetOrder,
		middleware.AuthMiddleware(&h.authClient))
	storeGroup.GET("/list", h.ListUserOrders,
		middleware.AuthMiddleware(&h.authClient))

}
