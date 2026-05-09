package httpserver

import (
	"gatewayapp/internal/client/authclient"
	"gatewayapp/internal/client/cartclient"
	"gatewayapp/internal/client/productclient"
	"gatewayapp/internal/client/storeclient"
	"gatewayapp/internal/client/userclient"
	"gatewayapp/internal/config"
	"gatewayapp/internal/delivery/http/carthanler"
	"gatewayapp/internal/delivery/http/producthandler"
	"gatewayapp/internal/delivery/http/storehandler"
	"gatewayapp/internal/delivery/http/userhandler"
	"myapp/pkg/logger"

	echoprom "github.com/labstack/echo-contrib/prometheus"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

type Server struct {
	config         config.Config
	storeHandler   storehandler.Handler
	userHandler    userhandler.Handler
	productHandler producthandler.Handler
	cartHandler    carthanler.Handler
	Router         *echo.Echo
}

func New(userClient userclient.Client,
	authClient authclient.Client,
	storeClient storeclient.Client,
	productClient productclient.Client,
	cartClient cartclient.Client,
	config config.Config) Server {
	return Server{
		Router:         echo.New(),
		userHandler:    userhandler.New(userClient, authClient),
		storeHandler:   storehandler.New(storeClient, authClient),
		productHandler: producthandler.New(productClient, authClient),
		cartHandler:    carthanler.New(cartClient, authClient),
		config:         config,
	}
}

func (s Server) Serve() {
	s.Router.Use(middleware.Recover())

	s.Router.Use(middleware.RequestID())

	s.Router.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:           true,
		LogStatus:        true,
		LogHost:          true,
		LogRemoteIP:      true,
		LogRequestID:     true,
		LogMethod:        true,
		LogContentLength: true,
		LogResponseSize:  true,
		LogLatency:       true,
		LogError:         true,
		LogProtocol:      true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			errMsg := ""
			if v.Error != nil {
				errMsg = v.Error.Error()
			}

			logger.Info("request",
				zap.String("request_id", v.RequestID),
				zap.String("host", v.Host),
				zap.String("content-length", v.ContentLength),
				zap.String("protocol", v.Protocol),
				zap.String("method", v.Method),
				zap.Duration("latency", v.Latency),
				zap.String("error", errMsg),
				zap.String("remote_ip", v.RemoteIP),
				zap.Int64("response_size", v.ResponseSize),
				zap.String("uri", v.URI),
				zap.Int("status", v.Status),
			)
			return nil
		},
	}))
	p := echoprom.NewPrometheus("gateway", nil)
	p.Use(s.Router)

	// Routes
	s.Router.GET("/health-check", s.healthCheck)
	s.userHandler.SetRoutes(s.Router)
	s.storeHandler.SetRoutes(s.Router)
	s.productHandler.SetRoutes(s.Router)
	s.cartHandler.SetRoutes(s.Router)

	// Start server
	s.Router.Logger.Fatal(s.Router.Start(s.config.HttpServer.Address))
}
