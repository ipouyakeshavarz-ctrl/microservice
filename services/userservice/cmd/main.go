package main

import (
	"fmt"
	"log"
	authpb "myapp/api/gen/auth"
	"myapp/pkg/config"
	cfg "userapp/internal/config"
	authclient "userapp/internal/delivery/grpc/auth"
	"userapp/internal/repository/migrator"
	"userapp/internal/repository/mysql"
	"userapp/internal/repository/mysql/mysqluser"
	"userapp/internal/validator"

	"userapp/internal/delivery/grpc"
	"userapp/internal/service"

	grpc2 "google.golang.org/grpc"
)

func main() {
	var cfg2 cfg.Config
	err := config.Load("config.yml", &cfg2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("cfg:%v\n", cfg2)

	mgr := migrator.New(cfg2.Mysql)
	mgr.Up()

	MysqlRepo := mysql.New(cfg2.Mysql)
	userRepo := mysqluser.New(MysqlRepo)

	conn, dErr := grpc2.Dial(fmt.Sprintf("127.0.0.1:%v", 50051), grpc2.WithInsecure())
	if dErr != nil {
		log.Fatalf("cannot connect to auth service: %v", dErr)
	}
	authClient := authpb.NewAuthServiceClient(conn)
	grpcAuthClient := authclient.NewGRPCAuthClient(authClient)

	userSvc := userservice.New(grpcAuthClient, userRepo)

	userV := validator.New(userRepo)

	grpcServer := grpc.NewServer(userV, userSvc, 50052)
	
	if err := grpcServer.Run(); err != nil {
		log.Fatal(err)
	}

}
