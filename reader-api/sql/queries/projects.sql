-- name: CreateProject :one
INSERT INTO projects (id, created_at, update_at, name, link)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetProjects :many
SELECT * FROM projects;