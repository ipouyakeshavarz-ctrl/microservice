package userhandler

import (
	"gatewayapp/internal/delivery/http/middleware"

	"github.com/labstack/echo/v4"
)

func (h Handler) SetRoutes(e *echo.Echo) {
	userGroup := e.Group("/users")

	userGroup.GET("/profile", h.Profile,
		middleware.AuthMiddleware(&h.authClient))
	userGroup.POST("/login", h.userLogin)
	userGroup.POST("/register", h.userRegister)
}
