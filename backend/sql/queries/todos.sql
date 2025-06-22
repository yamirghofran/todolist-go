-- name: GetTodos :many
SELECT * 
FROM todos
ORDER BY created_at DESC;

-- name: GetTodoByID :one
SELECT *
FROM todos
WHERE id = $1;

-- name: GetTodosByUserID :many
SELECT *
FROM todos
WHERE user_id = $1 AND id = $2
ORDER BY created_at DESC;

-- name: CreateTodo :one
INSERT INTO todos (user_id, title, description)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateTodo :one
UPDATE todos
SET title = $3, description = $4, is_completed = $5, updated_at = CURRENT_TIMESTAMP
WHERE user_id = $1 AND id = $2
RETURNING *;

-- name: DeleteTodo :exec
DELETE FROM todos WHERE user_id = $1 AND id = $2;

