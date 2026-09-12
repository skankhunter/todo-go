package tasks_service

import (
	"context"
	"fmt"

	"github.com/skankhunter/todo-go/internal/core/domain"
	core_errors "github.com/skankhunter/todo-go/internal/core/errors"
)

func (s *TasksSevice) GetTasks(
	ctx context.Context,
	userID *int,
	limit *int,
	offset *int,
) ([]domain.Task, error) {
	if limit != nil && *limit < 0 {
		return nil, fmt.Errorf(
			"limit must be non-negative: %w",
			core_errors.ErrnInvalidArgument,
		)
	}

	if offset != nil && *offset < 0 {
		return nil, fmt.Errorf(
			"offset must be non-negative: %w",
			core_errors.ErrnInvalidArgument,
		)
	}

	tasks, err := s.tasksRepository.GetTasks(ctx, userID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("get tasks from repository: %w", err)
	}

	return tasks, nil
}
