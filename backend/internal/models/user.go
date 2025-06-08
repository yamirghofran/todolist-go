package models

import (
	"errors"
	"time"
)

type User struct {
	ID           int32     `json:"id"`
	Name         string    `json:"title"`
	Email        string    `json:"description"`
	PasswordHash string    `json:"password_hash"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type CreateUserRequest struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

type UpdateUserRequest struct {
	Name         *string `json:"name"`
	Email        *string `json:"email"`
	PasswordHash *string `json:"password_hash"`
}

func (u *User) Validate() error {
	if u.Email == "" {
		return errors.New("Email is required")
	}
	if len(u.Name) > 255 {
		return errors.New("name must be less than 255 characters")
	}
	return nil
}
