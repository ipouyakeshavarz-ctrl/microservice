package carthanler

import (
	"gatewayapp/internal/delivery/http/middleware"

	"github.com/labstack/echo/v4"
)

func (h Handler) SetRoutes(e *echo.Echo) {
	storeGroup := e.Group("/cart")

	storeGroup.POST("/add_item", h.AddItem,
		middleware.AuthMiddleware(&h.authClient))
	storeGroup.GET("/check_out", h.Checkout,
		middleware.AuthMiddleware(&h.authClient))
	storeGroup.GET("/get_cart", h.GetCart,
		middleware.AuthMiddleware(&h.authClient))

}
