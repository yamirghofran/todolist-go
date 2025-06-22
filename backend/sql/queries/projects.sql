-- name: GetProjects :many
SELECT *
FROM projects
ORDER BY created_at DESC;

-- name: GetProjectByID :one
SELECT *
FROM projects
WHERE id = $1;

-- name: GetProjectsByUserID :many
SELECT *
FROM projects
WHERE user_id = $1
ORDER BY created_at DESC;

-- name: GetProjectAndTodosByUserID :many
SELECT sqlc.embed(projects), sqlc.embed(todos)
FROM projects
JOIN todos ON todos.project_id = projects.id
WHERE projects.user_id = $1 AND todos.user_id = $1 AND todos.project_id = $2;

-- name: GetProjectsAndTodosByUserID :many
SELECT sqlc.embed(projects), sqlc.embed(todos)
FROM projects
JOIN todos ON todos.project_id = projects.id
WHERE projects.user_id = $1 AND todos.user_id = $1;

-- name: CreateProject :one
INSERT INTO projects (user_id, title)
VALUES ($1, $2)
RETURNING *;

-- name: UpdateProject :one
UPDATE projects
SET title = $3, is_completed = $4, updated_at = CURRENT_TIMESTAMP
WHERE user_id = $1 AND id = $2
RETURNING *;

-- name: DeleteProject :exec
DELETE FROM projects
WHERE user_id = $1 AND id = $2;
