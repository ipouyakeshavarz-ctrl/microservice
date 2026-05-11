package main

import (
	"myapp/pkg/config"
	"myapp/pkg/logger"
	cfg "orderapp/internal/config"
	"orderapp/internal/delivery/broker/rabbitmq"
	"orderapp/internal/delivery/grpc"
	"orderapp/internal/repository/migrator"
	"orderapp/internal/repository/mysql"
	"orderapp/internal/repository/mysql/mysqlorder"
	orderservice "orderapp/internal/service"
	"os"
	"os/signal"
	"syscall"

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

	mgr := migrator.New(cfg2.Mysql)
	mgr.Up()

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

	err = ch.ExchangeDeclare(
		cfg2.RabbitMQ.Exchange,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		logger.Fatal("exchange declare failed", zap.Error(err))
	}

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

	err = ch.QueueBind(
		queue.Name,
		cfg2.RabbitMQ.CheckoutQueue,
		cfg2.RabbitMQ.Exchange,
		false,
		nil,
	)
	if err != nil {
		logger.Fatal("queue bind failed", zap.Error(err))
	}

	consumer := rabbitmq.New(ch, queue.Name, orderSvc)

	go func() {
		if err := consumer.Start(); err != nil {
			logger.Fatal("consumer start failed", zap.Error(err))
		}
	}()

	if err != nil {
		logger.Fatal("consumer start failed", zap.Error(err))
	}

	logger.Info("order service started and waiting for events")

	grpcServer := grpc.NewServer(orderSvc, cfg2.GrpcServer.OrderAddress, cfg2.Metrics.Port)

	go func() {
		logger.Info("🚀gRPC server started on ",
			zap.String("address", cfg2.GrpcServer.OrderAddress))

		if err := grpcServer.Run(); err != nil {
			logger.Fatal("cannot start grpc server", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	logger.Info("Received shutdown signal. Initiating graceful shutdown...")

	grpcServer.GracefulStop()

	if err := mysqlRepo.Conn().Close(); err != nil {
		logger.Error("failed to close MysqlRepo connection", zap.Error(err))
	}

	logger.Info("Graceful shutdown completed successfully. 🛑")
}
