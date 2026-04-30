package main

import (
	"fmt"
	"log"
	"myapp/pkg/config"
	authservice "userapp/internal/auth"
	cfg "userapp/internal/config"
	httpserver "userapp/internal/delivery/http"
	"userapp/internal/repository/migrator"
	"userapp/internal/repository/mysql"
	"userapp/internal/repository/mysql/mysqluser"
	userservice "userapp/internal/service"
	"userapp/internal/validator"
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
	authSvc := authservice.New(cfg2.Auth)
	userSvc := userservice.New(authSvc, userRepo)
	userV := validator.New(userRepo)
	server := httpserver.New(cfg2, userSvc, userV, authSvc)
	server.Serve()
}
