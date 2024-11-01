-- name: CreateTask :one
INSERT INTO tasks (id, title, status, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetTasks :many
SELECT * FROM tasks;

-- name: GetTasksWithProjects :many
SELECT
    tasks.*,
    json_build_object(
        'id', projects.id,
        'created_at', projects.created_at,
        'update_at', projects.update_at,
        'name', projects.name,
        'images', projects.images,
        'ocr_data', projects.ocr_data
    ) AS project
FROM tasks
LEFT JOIN projects ON projects.name = tasks.title;

-- name: UpdateTaskStatus :exec
UPDATE tasks
SET status = $1, updated_at = $2
WHERE id = $3;