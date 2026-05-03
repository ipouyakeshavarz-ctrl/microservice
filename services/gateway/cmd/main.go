package main

import (
	"fmt"
	"gatewayapp/internal/client/authclient"
	"gatewayapp/internal/client/storeclient"
	"gatewayapp/internal/client/userclient"
	httpserver "gatewayapp/internal/delivery/http"
	"log"
)

func main() {

	//cfg := config.Load()

	// gRPC connections
	//authConn, err := grpc.Dial(cfg.AuthServiceAddr, grpc.WithInsecure())
	//if err != nil {
	//	log.Fatal(err)
	//}
	fmt.Println("Hello world")
	// clients.

	authClient, err := authclient.New("127.0.0.1:50051")
	if err != nil {
		log.Fatal(err)
	}

	userClient, uErr := userclient.New("127.0.0.1:50052")
	if uErr != nil {
		log.Fatal(uErr)
	}

	storeClient, sErr := storeclient.New("127.0.0.1:50053")
	if sErr != nil {
		log.Fatal(uErr)
	}

	// echo
	fmt.Println("Hello world")

	server := httpserver.New(*userClient, *authClient, *storeClient)
	fmt.Println("Hello world")
	server.Serve()

}
