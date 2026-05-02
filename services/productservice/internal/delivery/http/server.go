package httpserver

import (
	"fmt"
	"productapp/internal/config"
	"productapp/internal/delivery/http/producthandler"
	productservice "productapp/internal/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	config         config.Config
	productHandler producthandler.Handler
	Router         *echo.Echo
}

func New(config config.Config, productSvc productservice.Service) Server {
	return Server{
		Router:         echo.New(),
		config:         config,
		productHandler: producthandler.New(productSvc),
	}
}

func (s Server) Serve() {
	s.Router.Use(middleware.RequestLogger())
	s.Router.Use(middleware.Recover())

	// Routes
	s.Router.GET("/health-check", s.healthCheck)
	s.productHandler.SetRoutes(s.Router)

	// Start server
	address := fmt.Sprintf(":%d", s.config.Httpserver.Port)
	fmt.Printf("start echo server on %s\n", address)
	if err := s.Router.Start(address); err != nil {
		fmt.Println("router start error", err)
	}
}
