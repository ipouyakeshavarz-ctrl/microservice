package producthandler

import (
	"gatewayapp/internal/delivery/http/middleware"

	"github.com/labstack/echo/v4"
)

func (h Handler) SetRoutes(e *echo.Echo) {
	storeGroup := e.Group("/product")

	storeGroup.POST("/create", h.createProduct,
		middleware.AuthMiddleware(&h.authClient))
	storeGroup.POST("/delete", h.deleteProduct,
		middleware.AuthMiddleware(&h.authClient))
	storeGroup.POST("/update", h.updateProduct,
		middleware.AuthMiddleware(&h.authClient))
	storeGroup.POST("/update", h.updateProduct,
		middleware.AuthMiddleware(&h.authClient))
	storeGroup.GET("/list", h.GetProduct,
		middleware.AuthMiddleware(&h.authClient))

}
