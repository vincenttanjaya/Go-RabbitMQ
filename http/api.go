package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vincenttanjaya/Go-RabbitMQ/model"
	"github.com/vincenttanjaya/Go-RabbitMQ/service"
)

type Handler struct {
	Service service.Usecase
}

func (h *Handler) Routes(s service.Usecase, e *echo.Echo) {

	h.Service = s

	// HealthCheck for checking server health
	e.GET("/", h.healthCheck)

	// GetAllTodo get all todo from todo:created queue and todo:deleted queue
	e.GET("/todo", h.getToDo)

	// [Testing] for testing publish create and delete
	publish := e.Group("/publish")
	publish.POST("/create", h.publishCreate)
	publish.POST("/delete", h.publishDelete)

}

// Handler for healthCheck
func (h *Handler) healthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
}

// Handler for publish to req:create:todo channel
func (h *Handler) publishCreate(c echo.Context) error {

	ctx := context.Background()
	var req model.PublishCreateRequest

	err := c.Bind(&req)
	if err != nil {
		fmt.Println("[publishCreate] failed to bind request ", req)
		return c.JSON(http.StatusBadRequest, model.PublishResponse{
			Message: "Failed to bind request",
		})
	}

	err = h.Service.PublishReqCreate(ctx, req.Name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.PublishResponse{
			Message: "Failed publish to channel req:create:todo",
		})
	}

	fmt.Println("[publishCreate] success publish to req:create:todo")
	return c.JSON(http.StatusOK, model.PublishResponse{
		Message: "Success publish to channel req:create:todo",
	})
}

// Handler for publish to req:create:delete channel
func (h *Handler) publishDelete(c echo.Context) error {

	ctx := context.Background()
	var req model.PublishDeleteRequest

	err := c.Bind(&req)
	if err != nil {
		fmt.Println("[publishDelete] failed to bind request ", req)
		return c.JSON(http.StatusBadRequest, model.PublishResponse{
			Message: "Failed to bind request",
		})
	}

	err = h.Service.PublishReqDelete(ctx, req.Uuid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.PublishResponse{
			Message: "Failed publish to channel req:create:delete",
		})
	}

	fmt.Println("[publishDelete] success publish to req:create:delete")
	return c.JSON(http.StatusOK, model.PublishResponse{
		Message: "Success publish to channel req:create:delete",
	})

}

// Handler for get all todo
func (h *Handler) getToDo(c echo.Context) error {
	ctx := context.Background()

	resp, err := h.Service.GetAllTodo(ctx)
	if err != nil {
		fmt.Println("[getToDo] failed to GetAllTodo")
		return c.JSON(http.StatusInternalServerError, "Failed to get todo")
	}

	return c.JSON(http.StatusOK, resp)
}
