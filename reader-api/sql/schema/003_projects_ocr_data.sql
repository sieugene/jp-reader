-- +goose Up
ALTER TABLE projects
    DROP COLUMN link,
    ADD COLUMN images TEXT[] NOT NULL,
    ADD COLUMN ocr_data JSONB NOT NULL;

-- +goose Down
ALTER TABLE projects
    ADD COLUMN link TEXT,
    DROP COLUMN images,
    DROP COLUMN ocr_data;
