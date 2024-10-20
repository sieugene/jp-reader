-- +goose Up
ALTER TABLE projects
    DROP COLUMN link,
    ADD COLUMN images TEXT[],
    ADD COLUMN ocrData JSONB;

-- +goose Down
ALTER TABLE projects
    ADD COLUMN link TEXT,
    DROP COLUMN images,
    DROP COLUMN ocrData;
