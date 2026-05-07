package main

import (
	"cartapp/internal/adapter/broker/rabbitmq"
	"cartapp/internal/adapter/redis"
	cfg "cartapp/internal/config"
	"cartapp/internal/delivery/grpc"
	"cartapp/internal/repository/redis/cartrepo"
	cartservice "cartapp/internal/service"
	"myapp/pkg/config"
	"myapp/pkg/logger"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	logger.Info("config", zap.Any("config", cfg2))

	redisAdapter := redis.NewAdapter(cfg2.Redis)

	cartTTL := time.Duration(cfg2.Redis.ProductTTLMinutes) * time.Minute
	cartRepo := cartrepo.New(redisAdapter, cartTTL)

	conn, dErr := amqp.Dial(cfg2.RabbitMQ.URL)
	if dErr != nil {
		panic("cant connect to rabbitmq: " + dErr.Error())
	}

	ch, cErr := conn.Channel()

	if cErr != nil {
		logger.Fatal("cant connection to rabbitMQ", zap.Error(cErr))
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
		logger.Fatal("Can't make Exchange", zap.Error(err))
	}

	publisher := rabbitmq.NewPublisher(ch, cfg2.RabbitMQ.Exchange, cfg2.RabbitMQ.CheckoutRoutingKey)

	cartService := cartservice.New(cartRepo, publisher)

	grpcServer := grpc.NewServer(*cartService, cfg2.GrpcServer.CartAddress)

	go func() {
		logger.Info("🚀gRPC server started on ",
			zap.String("address", cfg2.GrpcServer.CartAddress))

		if err := grpcServer.Run(); err != nil {
			logger.Fatal("cannot start grpc server", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	logger.Info("Received shutdown signal. Initiating graceful shutdown...")

	grpcServer.GracefulStop()

	if err := redisAdapter.Client().Close(); err != nil {
		logger.Error("failed to close Redis Adapter connection", zap.Error(err))
	}

	logger.Info("Graceful shutdown completed successfully. 🛑")

}
