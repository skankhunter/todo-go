package domain

import (
	"fmt"
	"time"

	core_errors "github.com/skankhunter/todo-go/internal/core/errors"
)

type Task struct {
	ID      int
	Version int

	Title       string
	Description *string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time

	AuthorUserId int
}

func NewTask(
	id int,
	version int,
	title string,
	description *string,
	completed bool,
	createdAt time.Time,
	completedAt *time.Time,
	authorUserId int,
) Task {
	return Task{
		ID:           id,
		Version:      version,
		Title:        title,
		Description:  description,
		Completed:    completed,
		CreatedAt:    createdAt,
		CompletedAt:  completedAt,
		AuthorUserId: authorUserId,
	}
}

func NewTaskUninitialized(
	title string,
	description *string,
	authorUserId int,
) Task {
	return NewTask(
		UnitializedID,
		UnitializedVersion,
		title,
		description,
		false,
		time.Now(),
		nil,
		authorUserId,
	)
}

func (t *Task) Validate() error {
	titleLen := len([]rune(t.Title))
	if titleLen < 1 || titleLen > 100 {
		return fmt.Errorf(
			"invalid title len: %d: %w",
			titleLen,
			core_errors.ErrnInvalidArgument,
		)
	}

	if t.Description != nil {
		descriptionLen := len([]rune(*t.Description))
		if descriptionLen < 1 || descriptionLen > 1000 {
			return fmt.Errorf(
				"invalid Description len: %d: %w",
				descriptionLen,
				core_errors.ErrnInvalidArgument,
			)
		}
	}

	if t.Completed {
		if t.CompletedAt == nil {
			return fmt.Errorf(
				"invalid CompletedAt cant be 'nil' if Completed 'true': %w",
				core_errors.ErrnInvalidArgument,
			)
		}

		if t.CompletedAt.Before(t.CreatedAt) {
			return fmt.Errorf(
				"invalid CompletedAt cant be before 'CreatedAt': %w",
				core_errors.ErrnInvalidArgument,
			)
		}
	} else {
		if t.CompletedAt != nil {
			return fmt.Errorf(
				"CompletedAt must be 'nil' if Completed 'false' : %w",
				core_errors.ErrnInvalidArgument,
			)
		}
	}
	return nil
}
