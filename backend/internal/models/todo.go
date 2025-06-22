package models

import (
	"errors"
	"time"
)

type Todo struct {
	ID          int32     `json:"id"`
	ProjectID   int32     `json:"project_id"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	IsCompleted bool      `json:"is_completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateTodoRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	ProjectID   int32  `json:"project_id"`
}

type UpdateTodoRequest struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	IsCompleted *bool   `json:"is_completed"`
	ProjectID   *int32  `json:"project_id"`
}

func (t *Todo) Validate() error {
	if t.Title == "" {
		return errors.New("title is required")
	}
	if len(t.Title) > 255 {
		return errors.New("title must be less than 255 characters")
	}
	return nil
}
