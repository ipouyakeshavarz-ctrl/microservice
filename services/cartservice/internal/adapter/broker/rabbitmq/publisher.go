package rabbitmq

import (
	"context"
	"encoding/json"

	"cartapp/internal/domain"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Publisher struct {
	ch         *amqp.Channel
	exchange   string
	RoutingKey string
}

func NewPublisher(ch *amqp.Channel, exchange string, RoutingKey string) *Publisher {
	return &Publisher{
		ch:         ch,
		exchange:   exchange,
		RoutingKey: RoutingKey,
	}
}

func (p *Publisher) PublishCartCheckedOut(ctx context.Context, event domain.CartCheckedOutEvent) error {

	body, err := json.Marshal(event)
	if err != nil {
		return err
	}

	return p.ch.PublishWithContext(
		ctx,
		p.exchange,
		p.RoutingKey,
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			DeliveryMode: amqp.Persistent,
			Body:         body,
		},
	)
}
