package tasks_service

import (
	"context"
	"fmt"
)

func (s *TasksSevice) DeleteTask(
	ctx context.Context,
	id int,
) error {
	if err := s.tasksRepository.DeleteTask(ctx, id); err != nil {
		return fmt.Errorf("delete task from repositor: %w", err)
	}
	return nil
}
