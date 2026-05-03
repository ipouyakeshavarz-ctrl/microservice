package httpserver

import (
	"fmt"
	"gatewayapp/internal/client/authclient"
	"gatewayapp/internal/client/productclient"
	"gatewayapp/internal/client/storeclient"
	"gatewayapp/internal/client/userclient"
	"gatewayapp/internal/config"
	"gatewayapp/internal/delivery/http/producthandler"
	"gatewayapp/internal/delivery/http/storehandler"
	"gatewayapp/internal/delivery/http/userhandler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	config         config.Config
	storeHandler   storehandler.Handler
	userHandler    userhandler.Handler
	productHandler producthandler.Handler
	Router         *echo.Echo
}

func New(userClient userclient.Client,
	authClient authclient.Client,
	storeClient storeclient.Client,
	productClient productclient.Client,
	config config.Config) Server {
	return Server{
		Router:         echo.New(),
		userHandler:    userhandler.New(userClient, authClient),
		storeHandler:   storehandler.New(storeClient, authClient),
		productHandler: producthandler.New(productClient, authClient),
		config:         config,
	}
}

func (s Server) Serve() {
	s.Router.Use(middleware.RequestLogger())
	s.Router.Use(middleware.Recover())

	// Routes
	s.Router.GET("/health-check", s.healthCheck)
	s.userHandler.SetRoutes(s.Router)
	s.storeHandler.SetRoutes(s.Router)
	s.productHandler.SetRoutes(s.Router)

	// Start server
	fmt.Printf("Listening on %s\n", s.config.HttpServer.Address)

	s.Router.Logger.Fatal(s.Router.Start(s.config.HttpServer.Address))
}
