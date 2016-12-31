
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE entries
ADD COLUMN language TEXT NOT NULL,
ADD COLUMN code TEXT NOT NULL,
ADD COLUMN exec_time DECIMAL;


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

ALTER TABLE entries
DROP COLUMN language,
DROP COLUMN code,
DROP COLUMN exec_time;
