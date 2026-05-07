package main

import (
	"myapp/pkg/config"
	"myapp/pkg/logger"
	cfg "orderapp/internal/config"
	"orderapp/internal/delivery/broker/rabbitmq"
	"orderapp/internal/repository/mysql"
	"orderapp/internal/repository/mysql/mysqlorder"
	orderservice "orderapp/internal/service"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

func main() {

	var cfg2 cfg.Config
	err := config.Load("config.yml", &cfg2)
	if err != nil {
		panic("failed to load config: " + err.Error())
	}

	logger.InitLogger(cfg2.Logger.ServiceName, cfg2.Logger.Development, cfg2.Logger.FilePath)
	defer logger.Sync()

	logger.Info("config loaded", zap.Any("config", cfg2))

	mysqlRepo := mysql.New(cfg2.Mysql)
	orderRepo := mysqlorder.New(mysqlRepo)

	orderSvc := orderservice.New(orderRepo)

	conn, err := amqp.Dial(cfg2.RabbitMQ.URL)
	if err != nil {
		logger.Fatal("cannot connect to rabbitmq", zap.Error(err))
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		logger.Fatal("cannot open rabbitmq channel", zap.Error(err))
	}
	defer ch.Close()

	queue, err := ch.QueueDeclare(
		"cart.checkout",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		logger.Fatal("queue declare failed", zap.Error(err))
	}

	consumer := rabbitmq.New(ch, queue.Name, orderSvc)

	err = consumer.Start()
	if err != nil {
		logger.Fatal("consumer start failed", zap.Error(err))
	}

	logger.Info("order service started and waiting for events")

	select {}
}
