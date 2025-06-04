package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yamirghofran/todolist-go/internal/models"
)

// Integration test - requires running PostgreSQL
func TestTodoService_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	// Setup test datanase connection
	service, err := NewTodoService("postgres://yamirghofran0:@localhost:5432/todogo_test?sslmode=disable")
	require.NoError(t, err)
	defer service.Close()

	ctx := context.Background()

	// Clean up before TestTodoService_Integration
	_, err = service.db.Exec(ctx, "TRUNCATE TABLE todos RESTART IDENTITY")
	require.NoError(t, err)

	t.Run("CreateTodo", func(t *testing.T) {
		req := models.CreateTodoRequest{
			Title:       "Test Todo",
			Description: "Test Description",
		}

		todo, err := service.CreateTodo(ctx, req)
		require.NoError(t, err)
		assert.Equal(t, req.Title, todo.Title)
		assert.Equal(t, req.Description, todo.Description)
		assert.False(t, todo.IsCompleted)
		assert.NotZero(t, todo.ID)
	})

	t.Run("GetTodos", func(t *testing.T) {
		todos, err := service.GetTodos(ctx)
		require.NoError(t, err)
		assert.Len(t, todos, 1)
	})

	t.Run("GetTodoByID", func(t *testing.T) {
		todo, err := service.GetTodoByID(ctx, 1)
		require.NoError(t, err)
		assert.Equal(t, "Test Todo", todo.Title)
	})

	t.Run("UpdateTodo", func(t *testing.T) {
		newTitle := "Updated Todo"
		isCompleted := true
		req := models.UpdateTodoRequest{
			Title:       &newTitle,
			IsCompleted: &isCompleted,
		}
		todo, err := service.UpdateTodo(ctx, 1, req)
		require.NoError(t, err)
		assert.Equal(t, newTitle, todo.Title)
		assert.True(t, todo.IsCompleted)
	})

	t.Run("DeleteTodo", func(t *testing.T) {
		err := service.DeleteTodo(ctx, 1)
		require.NoError(t, err)

		// Verify deletion
		_, err = service.GetTodoByID(ctx, 1)
		assert.Error(t, err)
	})
}
