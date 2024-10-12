-- +goose Up
CREATE TABLE projects (
	id UUID PRIMARY KEY,
	created_at TIMESTAMP NOT NULL,
	update_at TIMESTAMP NOT NULL,
	name TEXT NOT NULL,
	link TEXT NOT NULL
);

-- +goose Down
DROP TABLE projects;