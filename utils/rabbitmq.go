package utils

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"
	"github.com/vincenttanjaya/Go-RabbitMQ/constant"
	"github.com/vincenttanjaya/Go-RabbitMQ/model"
	"github.com/vincenttanjaya/Go-RabbitMQ/repository/rabbitmq"
	"github.com/vincenttanjaya/Go-RabbitMQ/service"
)

// Init rabbit mq connection from config
func NewRabbitMQConnetion(config RabbitMQConfig) (*amqp.Connection, error) {
	connection := fmt.Sprintf("amqp://%s:%s@%s:%s/", config.User, config.Password, config.Host, config.Port)
	fmt.Println(connection)
	return amqp.Dial(connection)
}

// A goroutine function for consume from req:create:todo queue and  req:delete:todo
func ReadConsume(repo rabbitmq.Repository, svc service.Usecase) {
	ctx := context.Background()
	createConsume, _ := repo.Consume(ctx, constant.ReqCreateTodoChannel)
	deleteConsume, _ := repo.Consume(ctx, constant.ReqDeleteTodoChannel)

	go func() {
		for d := range createConsume {
			message := model.Message{}
			json.Unmarshal(d.Body, &message)
			fmt.Println("Receive message ", message)
			svc.PublishTodoCreate(ctx, message.Name)
		}
	}()

	go func() {
		for d := range deleteConsume {
			message := model.Message{}
			json.Unmarshal(d.Body, &message)
			fmt.Println("Receive message ", message)
			svc.PublishTodoDelete(ctx, message.Id)
		}
	}()

}

// Declare Queue to avoid error when the connection haven't have the queue
func DeclareQueue(ch *amqp.Channel) {
	ch.QueueDeclare(
		constant.ReqCreateTodoChannel, // name
		false,                         // durable
		false,                         // delete when unused
		false,                         // exclusive
		false,                         // no-wait
		nil,                           // arguments
	)
	ch.QueueDeclare(
		constant.ReqDeleteTodoChannel, // name
		false,                         // durable
		false,                         // delete when unused
		false,                         // exclusive
		false,                         // no-wait
		nil,                           // arguments
	)

	ch.QueueDeclare(
		constant.TodoCreatedChannel, // name
		false,                       // durable
		false,                       // delete when unused
		false,                       // exclusive
		false,                       // no-wait
		nil,                         // arguments
	)

	ch.QueueDeclare(
		constant.TodoDeletedChannel, // name
		false,                       // durable
		false,                       // delete when unused
		false,                       // exclusive
		false,                       // no-wait
		nil,                         // arguments
	)
}
