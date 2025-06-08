package db

import (
	"context"
	"fmt"

	"github.com/yamirghofran/todolist-go/internal/models"
)

type TodoService struct {
	queries *Queries
}

func NewTodoService(queries *Queries) (*TodoService, error) {
	return &TodoService{
		queries: queries,
	}, nil
}

func (s *TodoService) CreateTodo(ctx context.Context, req models.CreateTodoRequest) (*models.Todo, error) {
	todo := &models.Todo{
		Title:       req.Title,
		Description: req.Description,
	}
	if err := todo.Validate(); err != nil {
		return nil, err
	}
	dbTodo, err := s.queries.CreateTodo(ctx, CreateTodoParams{
		Title:       req.Title,
		Description: &req.Description,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create todo: %w", err)
	}

	return &models.Todo{
		ID:          dbTodo.ID,
		Title:       dbTodo.Title,
		Description: *dbTodo.Description,
		IsCompleted: *dbTodo.IsCompleted,
		CreatedAt:   dbTodo.CreatedAt.Time,
		UpdatedAt:   dbTodo.UpdatedAt.Time,
	}, nil
}

func (s *TodoService) GetTodos(ctx context.Context) ([]models.Todo, error) {
	dbTodos, err := s.queries.GetTodos(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch todos: %w", err)
	}
	todos := make([]models.Todo, len(dbTodos))
	for _, dbTodo := range dbTodos {
		todo := models.Todo{
			Title:       dbTodo.Title,
			Description: *dbTodo.Description,
			IsCompleted: *dbTodo.IsCompleted,
			CreatedAt:   dbTodo.CreatedAt.Time,
			UpdatedAt:   dbTodo.UpdatedAt.Time,
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func (s *TodoService) GetTodoByID(ctx context.Context, id int32) (*models.Todo, error) {
	dbTodo, err := s.queries.GetTodoByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch todo by id: %w", err)
	}
	todo := &models.Todo{
		ID:          dbTodo.ID,
		Title:       dbTodo.Title,
		Description: *dbTodo.Description,
		IsCompleted: *dbTodo.IsCompleted,
		CreatedAt:   dbTodo.CreatedAt.Time,
		UpdatedAt:   dbTodo.CreatedAt.Time,
	}
	return todo, nil
}

func (s *TodoService) UpdateTodo(ctx context.Context, id int32, req models.UpdateTodoRequest) (*models.Todo, error) {
	existing, err := s.GetTodoByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch item to update: %w", err)
	}
	if req.Title != nil {
		existing.Title = *req.Title
	}
	if req.Description != nil {
		existing.Description = *req.Description
	}
	if req.IsCompleted != nil {
		existing.IsCompleted = *req.IsCompleted
	}
	if err := existing.Validate(); err != nil {
		return nil, fmt.Errorf("failed to validate todo: %w", err)
	}

	dbTodo, err := s.queries.UpdateTodo(ctx, UpdateTodoParams{
		ID:          id,
		Title:       existing.Title,
		Description: &existing.Description,
		IsCompleted: &existing.IsCompleted,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to update todo: %w", err)
	}
	return &models.Todo{
		ID:          dbTodo.ID,
		Title:       dbTodo.Title,
		Description: *dbTodo.Description,
		IsCompleted: *dbTodo.IsCompleted,
		CreatedAt:   dbTodo.CreatedAt.Time,
		UpdatedAt:   dbTodo.UpdatedAt.Time,
	}, nil
}

func (s *TodoService) DeleteTodo(ctx context.Context, id int32) error {
	err := s.queries.DeleteTodo(ctx, id)
	if err != nil {
		return fmt.Errorf("error deleting todo: %w", err)
	}
	return nil
}
