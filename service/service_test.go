package service

import (
	"context"
	"errors"
	"testing"

	assert "github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"github.com/vincenttanjaya/Go-RabbitMQ/constant"
	mock "github.com/vincenttanjaya/Go-RabbitMQ/mock/rabbitmq"
	"github.com/vincenttanjaya/Go-RabbitMQ/model"
)

func Test_PublishReqCreate(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mock.NewMockRepository(ctrl)
	svc := NewService(mockRepo)
	t.Run("success", func(t *testing.T) {
		ctx := context.Background()
		mockReq := model.PublishRequest{
			Message: model.Message{
				Name: "some name",
			},
			Channel: constant.ReqCreateTodoChannel,
		}

		mockRepo.EXPECT().Publish(ctx, mockReq).Return(nil)

		err := svc.PublishReqCreate(ctx, "some name")
		assert.Equal(t, err, nil)
	})

	t.Run("return error", func(t *testing.T) {
		ctx := context.Background()
		mockReq := model.PublishRequest{
			Message: model.Message{
				Name: "some name",
			},
			Channel: constant.ReqCreateTodoChannel,
		}

		mockRepo.EXPECT().Publish(ctx, mockReq).Return(errors.New("some error"))

		err := svc.PublishReqCreate(ctx, "some name")
		assert.Equal(t, err, errors.New("some error"))
	})
}

func Test_PublishReqDelete(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mock.NewMockRepository(ctrl)
	svc := NewService(mockRepo)
	t.Run("success", func(t *testing.T) {
		ctx := context.Background()
		mockReq := model.PublishRequest{
			Message: model.Message{
				Id: "some uuid",
			},
			Channel: constant.ReqDeleteTodoChannel,
		}

		mockRepo.EXPECT().Publish(ctx, mockReq).Return(nil)

		err := svc.PublishReqDelete(ctx, "some uuid")
		assert.Equal(t, err, nil)
	})

	t.Run("return error", func(t *testing.T) {
		ctx := context.Background()
		mockReq := model.PublishRequest{
			Message: model.Message{
				Id: "some uuid",
			},
			Channel: constant.ReqDeleteTodoChannel,
		}

		mockRepo.EXPECT().Publish(ctx, mockReq).Return(errors.New("some error"))

		err := svc.PublishReqDelete(ctx, "some uuid")
		assert.Equal(t, err, errors.New("some error"))
	})
}

func Test_PublishTodoCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mock.NewMockRepository(ctrl)
	svc := NewService(mockRepo)

	t.Run("success", func(t *testing.T) {
		ctx := context.Background()
		mockName := "some name"
		todoCreated = []model.Message{}
		mockRepo.EXPECT().Publish(ctx, gomock.Any()).Return(nil)

		err := svc.PublishTodoCreate(ctx, "some name")
		assert.Equal(t, err, nil)
		assert.Equal(t, todoCreated[0].Name, mockName)
	})

	t.Run("return error", func(t *testing.T) {
		ctx := context.Background()

		todoCreated = []model.Message{}
		mockRepo.EXPECT().Publish(ctx, gomock.Any()).Return(errors.New("some error"))

		err := svc.PublishTodoCreate(ctx, "some name")
		assert.Equal(t, err, errors.New("some error"))
		assert.Equal(t, todoCreated, []model.Message{})
	})

}

func Test_PublishTodoDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mock.NewMockRepository(ctrl)
	svc := NewService(mockRepo)

	t.Run("success", func(t *testing.T) {
		ctx := context.Background()
		mockUuid := "some uuid"
		todoCreated = []model.Message{
			{
				Id:   mockUuid,
				Name: "some name",
			},
		}
		todoDeleted = []model.Message{}
		mockRepo.EXPECT().Publish(ctx, gomock.Any()).Return(nil)

		err := svc.PublishTodoDelete(ctx, mockUuid)
		assert.Equal(t, err, nil)
		assert.Equal(t, todoDeleted[0].Id, mockUuid)
		assert.Equal(t, todoDeleted[0].Name, "some name")
	})

	t.Run("return error", func(t *testing.T) {
		ctx := context.Background()
		mockUuid := "some uuid"
		todoCreated = []model.Message{
			{
				Id:   mockUuid,
				Name: "some name",
			},
		}
		todoDeleted = []model.Message{}
		mockRepo.EXPECT().Publish(ctx, gomock.Any()).Return(errors.New("some error"))

		err := svc.PublishTodoDelete(ctx, mockUuid)
		assert.Equal(t, err, errors.New("some error"))
		assert.Equal(t, todoDeleted, []model.Message{})
	})
}

func Test_GetAllTodo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockRepo := mock.NewMockRepository(ctrl)
	svc := NewService(mockRepo)

	t.Run("success", func(t *testing.T) {
		todoCreated = []model.Message{
			{
				Id:      "created 1",
				Name:    "name 1",
				Created: 12345,
			},
			{
				Id:      "created 2",
				Name:    "name 2",
				Created: 23456,
			},
		}

		todoDeleted = []model.Message{
			{
				Id:      "deleted 1",
				Name:    "name 1",
				Created: 12345,
			},
			{
				Id:      "deleted 2",
				Name:    "name 2",
				Created: 23456,
			},
		}
		ctx := context.Background()
		resp, err := svc.GetAllTodo(ctx)
		assert.Equal(t, err, nil)
		assert.Equal(t, resp.Created, todoCreated)
		assert.Equal(t, resp.Deleted, todoDeleted)
	})
}
