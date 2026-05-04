-- +goose Up
ALTER TABLE events ADD COLUMN user_id UUID NOT NULL REFERENCES users(id);

-- +goose Down
ALTER TABLE events DROP COLUMN user_id;
