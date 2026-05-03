package storehandler

import (
	"storeapp/internal/delivery/http/middleware"

	"github.com/labstack/echo/v4"
)

func (h Handler) SetRoutes(e *echo.Echo) {
	storeGroup := e.Group("/store")

	storeGroup.POST("/create", h.createStore)
	storeGroup.POST("/delete", h.deleteStore,
		middleware.Auth(h.authSvc, h.authConfig))
	storeGroup.POST("/update", h.updateStore,
		middleware.Auth(h.authSvc, h.authConfig))
	storeGroup.POST("/update", h.updateStore,
		middleware.Auth(h.authSvc, h.authConfig))
	storeGroup.GET("/list", h.listStore,
		middleware.Auth(h.authSvc, h.authConfig))

}
