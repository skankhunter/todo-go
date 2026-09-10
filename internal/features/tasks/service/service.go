package tasks_service

import (
	"context"

	"github.com/skankhunter/todo-go/internal/core/domain"
)

type TasksSevice struct {
	tasksRepository TasksRepository
}

type TasksRepository interface {
	CreateTask(
		ctx context.Context,
		task domain.Task,
	) (domain.Task, error)
}

func NewTasksService(
	tasksRepository TasksRepository,
) *TasksSevice {
	return &TasksSevice{
		tasksRepository: tasksRepository,
	}
}
