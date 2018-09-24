
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE tasks ADD note text;


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE tasks DROP note;
