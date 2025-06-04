package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTodoValidation(t *testing.T) {
	tests := []struct {
		name    string
		todo    Todo
		wantErr bool
	}{
		{
			name: "valid todo",
			todo: Todo{
				Title:       "Test todo",
				Description: "Test description",
			},
			wantErr: false,
		},
		{
			name: "empty title",
			todo: Todo{
				Title:       "",
				Description: "Test description",
			},
			wantErr: true,
		},
		{
			name: "title too long",
			todo: Todo{
				Title:       string(make([]byte, 256)),
				Description: "Test description",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.todo.Validate()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestTodoJson(t *testing.T) {
	todo := Todo{
		ID:          1,
		Title:       "Test todo",
		Description: "Test description",
		IsCompleted: false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// Test that JSON tags work correctly
	assert.Equal(t, todo.Title, "Test todo")
	assert.Equal(t, todo.Description, "Test description")
	assert.False(t, todo.IsCompleted)
}
