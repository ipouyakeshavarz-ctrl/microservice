package router

import (
	"github.com/labstack/echo/v4"

	"gatewayapp/internal/handler"
	"gatewayapp/internal/middleware"
)

func Register(
	e *echo.Echo,
	auth handler.AuthClient,
	// user handler.UserClient,
	// product handler.ProductClient,
) {

	authHandler := handler.NewAuthHandler(auth)
	//userHandler := handler.NewUserHandler(user)
	//productHandler := handler.NewProductHandler(product)

	authMiddleware := middleware.NewAuthMiddleware(auth)

	api := e.Group("/api")

	// auth
	api.POST("/login", authHandler.Login)

	// user
	userGroup := api.Group("/users")
	userGroup.Use(authMiddleware)
	userGroup.GET("/profile", userHandler.Profile)

	// product
	productGroup := api.Group("/products")
	productGroup.Use(authMiddleware)

	productGroup.GET("", productHandler.List)
	productGroup.POST("", productHandler.Create)
}
