package rabbitmq

import (
	"context"
	"encoding/json"
	"log"
	"orderapp/internal/param"

	amqp "github.com/rabbitmq/amqp091-go"

	"orderapp/internal/domain"
	"orderapp/internal/service"
)

type Consumer struct {
	channel *amqp.Channel
	queue   string
	svc     *orderservice.Service
}

func New(channel *amqp.Channel, queue string, svc *orderservice.Service) *Consumer {
	return &Consumer{
		channel: channel,
		queue:   queue,
		svc:     svc,
	}
}

func (c *Consumer) Start() error {

	msgs, err := c.channel.Consume(
		c.queue,
		"",
		true, // auto ack
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	go func() {
		for msg := range msgs {

			var event domain.CartCheckedOutEvent

			err := json.Unmarshal(msg.Body, &event)
			if err != nil {
				log.Println("failed to unmarshal event:", err)
				continue
			}

			err = c.svc.CreateFromCheckout(context.Background(),
				param.CreateFromCheckoutRequest{Event: event})
			if err != nil {
				log.Println("failed to create order:", err)
				continue
			}

			log.Println("order created for user:", event.UserID)
		}
	}()

	return nil
}
