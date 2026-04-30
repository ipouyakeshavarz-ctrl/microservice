package httpserver

import (
	"fmt"
	authservice "storeapp/internal/auth"
	"storeapp/internal/config"
	"storeapp/internal/delivery/http/storehandler"
	storeservice "storeapp/internal/service"
	storevalidator "storeapp/internal/validator"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	config       config.Config
	storeHandler storehandler.Handler
	Router       *echo.Echo
}

func New(config config.Config, authSvc authservice.Service, storeSvc storeservice.Service,
	storeV storevalidator.Validator) Server {
	return Server{
		Router:       echo.New(),
		config:       config,
		storeHandler: storehandler.New(config.Auth, authSvc, storeSvc, storeV),
	}
}

func (s Server) Serve() {
	s.Router.Use(middleware.RequestLogger())
	s.Router.Use(middleware.Recover())

	// Routes
	s.Router.GET("/health-check", s.healthCheck)
	s.storeHandler.SetRoutes(s.Router)

	// Start server
	address := fmt.Sprintf(":%d", s.config.Httpserver.Port)
	fmt.Printf("start echo server on %s\n", address)
	if err := s.Router.Start(address); err != nil {
		fmt.Println("router start error", err)
	}
}
