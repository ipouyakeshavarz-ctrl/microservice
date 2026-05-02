package producthandler

import (
	"github.com/labstack/echo/v4"
)

func (h Handler) SetRoutes(e *echo.Echo) {
	storeGroup := e.Group("/product")

	storeGroup.POST("/create", h.createProduct)
	storeGroup.POST("/delete", h.deleteProduct)
	storeGroup.POST("/update", h.updateProduct)
	storeGroup.POST("/update", h.updateProduct)
	storeGroup.GET("/list", h.GetProduct)

}
