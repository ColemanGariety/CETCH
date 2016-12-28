
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

ALTER TABLE competitions
ADD COLUMN date DATE,
ADD COLUMN solution DECIMAL;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

ALTER TABLE competitions
DROP COLUMN date,
DROP COLUMN solution;
