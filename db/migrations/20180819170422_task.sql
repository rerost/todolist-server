
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE tasks (
  id integer,
  title text,
  created_at timestamp
);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE tasks;
