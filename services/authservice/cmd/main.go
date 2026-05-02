package main

import (
	cfg "authapp/internal/config"
	authservice "authapp/internal/service"
	"fmt"
	"log"
	gen "myapp/api/gen/auth"
	"myapp/pkg/config"
	"net"

	"google.golang.org/grpc"
)

func main() {
	var cfg2 cfg.Config
	err := config.Load("config.yml", &cfg2)
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	authServer := authservice.New(cfg2.Auth)

	gen.RegisterAuthServiceServer(grpcServer, authServer)

	fmt.Println("🚀 AuthService gRPC server running on port 50051")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
