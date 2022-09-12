package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"
	"github.com/vincenttanjaya/Go-RabbitMQ/model"
)

type RabbitMQRepo struct {
	ch *amqp.Channel
}

func NewRabbitMQRepo(ch *amqp.Channel) Repository {
	return &RabbitMQRepo{
		ch: ch,
	}
}

// Publish implements publish message to specific channel
func (p *RabbitMQRepo) Publish(ctx context.Context, req model.PublishRequest) error {
	msg, err := json.Marshal(req.Message)
	if err != nil {
		fmt.Println("[Repo:Publish] failed to marshal message")
		return err
	}

	err = p.ch.Publish(
		"",
		req.Channel,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        msg,
		},
	)

	if err != nil {
		fmt.Println("[Repo:Publish] failed to publish message to channel ", req.Channel, " error: ", err.Error())
		return err
	}

	return nil
}
