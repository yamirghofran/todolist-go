-- name: GetTodos :many
SELECT id, title, description, is_completed, created_at, updated_at
FROM todos
ORDER BY created_at DESC;

-- name: GetTodoByID :one
SELECT id, title, description, is_completed, created_at, updated_at
FROM todos
WHERE id = $1;

-- name: CreateTodo :one
INSERT INTO todos (title, description)
VALUES ($1, $2)
RETURNING id, title, description, is_completed, created_at, updated_at;

-- name: UpdateTodo :one
UPDATE todos
SET title = $2, description = $3, is_completed = $4, updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING id, title, description, is_completed, created_at, updated_at;

-- name: DeleteTodo :exec
DELETE FROM todos WHERE id = $1;

