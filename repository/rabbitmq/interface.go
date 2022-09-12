//go:generate mockgen -destination=../../mock/rabbitmq/rabbitmq.go github.com/vincenttanjaya/Go-RabbitMQ/repository/rabbitmq Repository
package rabbitmq

import (
	"context"

	"github.com/streadway/amqp"
	"github.com/vincenttanjaya/Go-RabbitMQ/model"
)

type Repository interface {
	// Publish implements publish message to specific channel
	Publish(ctx context.Context, req model.PublishRequest) error
	// Consume implements consume from specific channel
	Consume(ctx context.Context, channel string) (<-chan amqp.Delivery, error)
}
