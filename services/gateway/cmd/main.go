package main

import (
	"gatewayapp/internal/client/authclient"
	"gatewayapp/internal/client/productclient"
	"gatewayapp/internal/client/storeclient"
	"gatewayapp/internal/client/userclient"
	cfg "gatewayapp/internal/config"
	httpserver "gatewayapp/internal/delivery/http"
	"log"
	"myapp/pkg/config"
)

func main() {

	var cfg2 cfg.Config
	err := config.Load("config.yml", &cfg2)

	authClient, err := authclient.New(cfg2.GrpcClient.ProductAddress)
	if err != nil {
		log.Fatal(err)
	}

	userClient, uErr := userclient.New(cfg2.GrpcClient.UserAddress)
	if uErr != nil {
		log.Fatal(uErr)
	}

	storeClient, sErr := storeclient.New(cfg2.GrpcClient.StoreAddress)
	if sErr != nil {
		log.Fatal(uErr)
	}

	productClient, pErr := productclient.New(cfg2.GrpcClient.ProductAddress)
	if pErr != nil {
		log.Fatal(pErr)
	}

	// echo

	server := httpserver.New(*userClient, *authClient, *storeClient, *productClient, cfg2)

	server.Serve()

}
