package storehandler

import (
	"gatewayapp/internal/delivery/http/middleware"

	"github.com/labstack/echo/v4"
)

func (h Handler) SetRoutes(e *echo.Echo) {
	storeGroup := e.Group("/store")

	storeGroup.POST("/create", h.createStore,
		middleware.AuthMiddleware(&h.authClient))
	storeGroup.POST("/delete", h.deleteStore,
		middleware.AuthMiddleware(&h.authClient))
	storeGroup.POST("/update", h.updateStore,
		middleware.AuthMiddleware(&h.authClient))
	storeGroup.GET("/list", h.listStore,
		middleware.AuthMiddleware(&h.authClient))

}
