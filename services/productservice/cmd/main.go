package main

import (
	"fmt"
	"log"
	"myapp/pkg/config"
	cfg "productapp/internal/config"
	httpserver "productapp/internal/delivery/http"
	"productapp/internal/repository/migrator"
	"productapp/internal/repository/mysql"
	mysqlproduct "productapp/internal/repository/mysql/mysqlproduct"
	productservice "productapp/internal/service"
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
	productRepo := mysqlproduct.New(MysqlRepo)

	productSvc := productservice.New(productRepo)

	server := httpserver.New(cfg2, productSvc)

	server.Serve()
}
