
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE INDEX user_competition_entry_index ON entries (user_id, competition_id);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP INDEX user_competition_entry_index;
