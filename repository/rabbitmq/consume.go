package rabbitmq

import (
	"context"

	"github.com/streadway/amqp"
)

// Consume implements consume from specific channel
func (p *RabbitMQRepo) Consume(ctx context.Context, channel string) (<-chan amqp.Delivery, error) {
	msgs, err := p.ch.Consume(
		channel,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	return msgs, err
}
