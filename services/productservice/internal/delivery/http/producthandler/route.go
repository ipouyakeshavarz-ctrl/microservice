package producthandler

import (
	"productapp/internal/delivery/http/middleware"

	"github.com/labstack/echo/v4"
)

func (h Handler) SetRoutes(e *echo.Echo) {
	storeGroup := e.Group("/product")

	storeGroup.POST("/create", h.createProduct,
		middleware.Auth(h.authSvc, h.authConfig))
	storeGroup.POST("/delete", h.deleteProduct,
		middleware.Auth(h.authSvc, h.authConfig))
	storeGroup.POST("/update", h.updateProduct,
		middleware.Auth(h.authSvc, h.authConfig))
	storeGroup.POST("/update", h.updateProduct,
		middleware.Auth(h.authSvc, h.authConfig))
	storeGroup.GET("/list", h.GetProduct,
		middleware.Auth(h.authSvc, h.authConfig))

}
