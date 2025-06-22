package db

import (
	"context"
	"fmt"

	"github.com/yamirghofran/todolist-go/internal/models"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	queries *Queries
}

func NewUserService(queries *Queries) *UserService {
	return &UserService{
		queries: queries,
	}
}

func (s *UserService) CreateUser(ctx context.Context, req models.CreateUserRequest) (*models.User, error) {
	var hashedPassword *string
	if req.Password != "" {
		hashedBytes, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
		if err != nil {
			return nil, fmt.Errorf("failed to hash password %w", err)
		}
		h := string(hashedBytes)
		hashedPassword = &h
	}
	user := &models.User{
		Name:           req.Name,
		Email:          req.Email,
		HashedPassword: *hashedPassword,
	}
	if err := user.Validate(); err != nil {
		return nil, err
	}
	dbUser, err := s.queries.CreateUser(ctx, CreateUserParams{
		Name:           &req.Name,
		Email:          req.Email,
		HashedPassword: hashedPassword,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return &models.User{
		ID:             dbUser.ID,
		Name:           *dbUser.Name,
		Email:          dbUser.Email,
		HashedPassword: *dbUser.HashedPassword,
	}, nil
}

func (s *UserService)