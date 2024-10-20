-- name: CreateProject :one
INSERT INTO projects (id, created_at, update_at, name, images, ocr_data)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetProjects :many
SELECT * FROM projects;

-- name: DeleteProjectByName :exec
DELETE FROM projects WHERE name = $1;