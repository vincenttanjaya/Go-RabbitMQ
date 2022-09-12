package constant

import "time"

const (
	// RabbitMQ Channel
	ReqCreateTodoChannel = "req:create:todo"
	ReqDeleteTodoChannel = "req:delete:todo"
	TodoCreatedChannel   = "todo:created"
	TodoDeletedChannel   = "todo:deleted"
)

var (
	TimeLocationAsiaJakarta, _ = time.LoadLocation("Asia/Jakarta")
)
