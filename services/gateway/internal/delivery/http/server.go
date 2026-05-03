package httpserver

import (
	"fmt"
	"gatewayapp/internal/client/authclient"
	"gatewayapp/internal/client/userclient"
	"gatewayapp/internal/delivery/http/userhandler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	userHandler userhandler.Handler
	Router      *echo.Echo
}

func New(userClient userclient.Client, authClient authclient.Client) Server {
	return Server{
		Router:      echo.New(),
		userHandler: userhandler.New(userClient, authClient),
	}
}

func (s Server) Serve() {
	s.Router.Use(middleware.RequestLogger())
	s.Router.Use(middleware.Recover())

	// Routes
	s.Router.GET("/health-check", s.healthCheck)
	s.userHandler.SetRoutes(s.Router)

	// Start server

	fmt.Printf("start echo server on %s\n", "127.0.0.1:8887")
	if err := s.Router.Start("127.0.0.1:8887"); err != nil {
		fmt.Println("router start error", err)
	}
}
