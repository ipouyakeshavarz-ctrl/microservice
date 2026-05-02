package main

import (
	"fmt"
	"log"
	authpb "myapp/api/gen/auth"
	"myapp/pkg/config"
	cfg "userapp/internal/config"
	authclient "userapp/internal/delivery/grpc"
	httpserver "userapp/internal/delivery/http"
	"userapp/internal/repository/migrator"
	"userapp/internal/repository/mysql"
	"userapp/internal/repository/mysql/mysqluser"
	userservice "userapp/internal/service"
	"userapp/internal/validator"

	"google.golang.org/grpc"
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

	conn, _ := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	authClient := authpb.NewAuthServiceClient(conn)
	grpcAuthClient := authclient.NewGRPCAuthClient(authClient)

	MysqlRepo := mysql.New(cfg2.Mysql)
	userRepo := mysqluser.New(MysqlRepo)
	userSvc := userservice.New(grpcAuthClient, userRepo)

	userV := validator.New(userRepo)

	server := httpserver.New(cfg2, userSvc, userV)
	server.Serve()
}
