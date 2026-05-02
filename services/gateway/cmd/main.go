package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"

	"gatewayapp/internal/client/authgrpc"
	"gatewayapp/internal/config"
	"gatewayapp/internal/router"
)

func main() {

	cfg := config.Load()

	// gRPC connections
	authConn, err := grpc.Dial(cfg.AuthServiceAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	userConn, err := grpc.Dial(cfg.UserServiceAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	productConn, err := grpc.Dial(cfg.ProductServiceAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	// clients
	authClient := authgrpc.NewAuthClient(authConn)
	userClient := grpcclient.NewUserClient(userConn)
	productClient := grpcclient.NewProductClient(productConn)

	// echo
	e := echo.New()

	router.Register(e, authClient)

	log.Println("gateway running on", cfg.Port)

	e.Logger.Fatal(e.Start(cfg.Port))
}
