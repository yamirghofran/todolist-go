package models

import (
	"errors"
	"time"
)

type User struct {
	ID             int32     `json:"id"`
	Name           string    `json:"name,omitempty"`
	Email          string    `json:"email"`
	HashedPassword string    `json:"-"`
	OAuthProvider  string    `json:"oauth_provider,omitempty"`
	OAuthID        string    `json:"oauth_id,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type CreateUserRequest struct {
	Name          string `json:"name"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	OAuthProvider string `json:"oauth_provider"`
	OAuthID       string `json:"oauth_id"`
}

type UpdateUserRequest struct {
	Name          *string `json:"name"`
	Email         *string `json:"email"`
	Password      *string `json:"password"`
	OAuthProvider *string `json:"oauth_provider"`
	OAuthID       *string `json:"oauth_id"`
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
