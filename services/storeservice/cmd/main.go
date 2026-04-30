package main

import (
	"fmt"
	"log"
	"myapp/pkg/config"
	authservice "storeapp/internal/auth"
	cfg "storeapp/internal/config"
	httpserver "storeapp/internal/delivery/http"
	"storeapp/internal/repository/migrator"
	"storeapp/internal/repository/mysql"
	"storeapp/internal/repository/mysql/mysqlstore"
	storeservice "storeapp/internal/service"
	storevalidator "storeapp/internal/validator"
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
	storeRepo := mysqlstore.New(MysqlRepo)
	authSvc := authservice.New(cfg2.Auth)
	storeSvc := storeservice.New(storeRepo)
	storeV := storevalidator.New()
	server := httpserver.New(cfg2, authSvc, storeSvc, storeV)
	server.Serve()
}
