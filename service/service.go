package service

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/vincenttanjaya/Go-RabbitMQ/constant"
	"github.com/vincenttanjaya/Go-RabbitMQ/model"
	"github.com/vincenttanjaya/Go-RabbitMQ/repository/rabbitmq"
)

type Usecase interface {
	PublishReqCreate(ctx context.Context, msg string) error
	PublishReqDelete(ctx context.Context, uuid string) error
	PublishTodoCreate(ctx context.Context, msg string) error
	PublishTodoDelete(ctx context.Context, uuid string) error
	GetAllTodo(ctx context.Context) (model.GetAllTodoResponse, error)
}

type Service struct {
	PublishRepo rabbitmq.Repository
}

// local cache for save todo created data
var todoCreated = []model.Message{}

// local cache for save todo delete data
var todoDeleted = []model.Message{}

func NewService(p rabbitmq.Repository) Usecase {
	return &Service{
		PublishRepo: p,
	}
}

// Service for publish to req:create:todo queue (for testing)
func (s *Service) PublishReqCreate(ctx context.Context, msg string) error {

	var req model.PublishRequest
	req.Message = model.Message{
		Name: msg,
	}
	req.Channel = constant.ReqCreateTodoChannel

	err := s.PublishRepo.Publish(ctx, req)
	if err != nil {
		fmt.Println("[Service:PublishReqCreate] failed to publish ", err.Error())
		return err
	}

	return nil
}

// Service for publish to req:create:delete queue (for testing)
func (s *Service) PublishReqDelete(ctx context.Context, uuid string) error {
	var req model.PublishRequest
	req.Message = model.Message{
		Id: uuid,
	}
	req.Channel = constant.ReqDeleteTodoChannel

	err := s.PublishRepo.Publish(ctx, req)
	if err != nil {
		fmt.Println("[Service:PublishReqDelete] failed to publish ", err.Error())
		return err
	}

	return nil
}

// Service for publish to todo:created queue (triggered when consume from req:create:todo)
func (s *Service) PublishTodoCreate(ctx context.Context, msg string) error {
	var req model.PublishRequest
	req.Message = model.Message{
		Id:      uuid.NewString(),
		Name:    msg,
		Created: time.Now().In(constant.TimeLocationAsiaJakarta).Unix(),
	}
	req.Channel = constant.TodoCreatedChannel

	err := s.PublishRepo.Publish(ctx, req)
	if err != nil {
		fmt.Println("[Service:PublishTodoCreate] failed to publish. Error: ", err.Error())
		return err
	}

	// if success save to local cache
	todoCreated = append(todoCreated, req.Message)

	return nil
}

// Service for publish to todo:deleted queue (triggered when consume from req:delete:todo)
func (s *Service) PublishTodoDelete(ctx context.Context, uuid string) error {
	var req model.PublishRequest
	var name string
	for _, todo := range todoCreated {
		if todo.Id == uuid {
			name = todo.Name
			break
		}
	}

	req.Message = model.Message{
		Id:      uuid,
		Name:    name,
		Created: time.Now().In(constant.TimeLocationAsiaJakarta).Unix(),
	}
	req.Channel = constant.TodoDeletedChannel

	err := s.PublishRepo.Publish(ctx, req)
	if err != nil {
		fmt.Println("[Service:PublishTodoDeleted] failed to publish. Error: ", err.Error())
		return err
	}

	// if success save to local cache
	todoDeleted = append(todoDeleted, req.Message)

	return nil
}

// Service for get all todo from local cache
func (s *Service) GetAllTodo(ctx context.Context) (model.GetAllTodoResponse, error) {
	resp := model.GetAllTodoResponse{}

	// get response from local cache
	resp.Created = todoCreated
	resp.Deleted = todoDeleted

	return resp, nil
}
