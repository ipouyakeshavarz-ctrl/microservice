package httpserver

import (
	"fmt"
	authservice "userapp/internal/auth"
	"userapp/internal/config"

	"userapp/internal/service"

	"userapp/internal/delivery/http/userhandler"
	"userapp/internal/validator"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	config      config.Config
	userHandler userhandler.Handler
	Router      *echo.Echo
}

func New(config config.Config, userSvc userservice.Service,
	userValidator validator.Validator, authSvc authservice.Service) Server {
	return Server{
		Router:      echo.New(),
		config:      config,
		userHandler: userhandler.New(userSvc, userValidator, config.Auth, authSvc),
	}
}

func (s Server) Serve() {
	s.Router.Use(middleware.RequestLogger())
	s.Router.Use(middleware.Recover())

	// Routes
	s.Router.GET("/health-check", s.healthCheck)
	s.userHandler.SetRoutes(s.Router)

	// Start server
	address := fmt.Sprintf(":%d", s.config.Httpserver.Port)
	fmt.Printf("start echo server on %s\n", address)
	if err := s.Router.Start(address); err != nil {
		fmt.Println("router start error", err)
	}
}
