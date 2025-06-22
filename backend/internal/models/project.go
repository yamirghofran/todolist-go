package models

import "time"

type Project struct {
	ID          int32     `json:"id"`
	UserID      int32     `json:"user_id"`
	Title       string    `json:"title"`
	IsCompleted *bool     `json:"is_completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateProjectRequest struct {
	Title string `json:"title"`
}

type UpdateProjectRequest struct {
	Title       *string `json:"title"`
	IsCompleted bool    `json:"is_completed"`
}
