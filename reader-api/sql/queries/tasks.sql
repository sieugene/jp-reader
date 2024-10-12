-- name: CreateTask :one
INSERT INTO tasks (id, title, status, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetTasks :many
SELECT * FROM tasks;

-- name: UpdateTaskStatus :exec
UPDATE tasks
SET status = $1, updated_at = $2
WHERE id = $3;