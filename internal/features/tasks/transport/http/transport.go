package tasks_transport_http

import (
	"context"
	"net/http"

	"github.com/skankhunter/todo-go/internal/core/domain"
	core_http_server "github.com/skankhunter/todo-go/internal/core/transport/http/server"
)

type TasksHTTPHandler struct {
	taskService TasksSevice
}

type TasksSevice interface {
	CreateTask(
		ctx context.Context,
		task domain.Task,
	) (domain.Task, error)
}

func NewTasksHTTPHandler(
	taskService TasksSevice,
) *TasksHTTPHandler {
	return &TasksHTTPHandler{
		taskService: taskService,
	}
}

func (h *TasksHTTPHandler) Routes() []core_http_server.Route {
	return []core_http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/tasks",
			Handler: h.CreateTask,
		},
	}
}
