package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/vincenttanjaya/Go-RabbitMQ/http"
	"github.com/vincenttanjaya/Go-RabbitMQ/repository/rabbitmq"
	"github.com/vincenttanjaya/Go-RabbitMQ/service"
	"github.com/vincenttanjaya/Go-RabbitMQ/utils"
)

func main() {

	// Read config from config.yaml
	cfg, err := utils.InitServiceConfig()
	if err != nil {
		fmt.Println(err)
		panic("invalid config.yaml")
	}

	// create new rabbitMQ connection and channel
	conn, err := utils.NewRabbitMQConnetion(cfg.RabbitMQConfig)
	if err != nil {
		fmt.Println(err)
		panic("invalid rabbitmq connection")
	}
	ch, err := conn.Channel()

	// this is for declare queue in the connection with 4 queue -> req:create:todo, req:delete:todo, todo:created, todo:deleted
	utils.DeclareQueue(ch)

	// create a rabbitMQ repository
	rabbitmqRepo := rabbitmq.NewRabbitMQRepo(ch)
	svc := service.NewService(rabbitmqRepo)

	fmt.Println("success connect to rabbitmq")

	// create a http request handler
	route := echo.New()
	httpHandler := &http.Handler{}
	httpHandler.Routes(svc, route)

	quit := make(chan os.Signal)

	// consume go routine
	utils.ReadConsume(rabbitmqRepo, svc)

	route.Start(cfg.ServerConfig.Port)
	<-quit
	fmt.Println("Bye !")
}
