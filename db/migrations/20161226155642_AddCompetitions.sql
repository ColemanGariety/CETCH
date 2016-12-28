
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE competitions (
       id SERIAL UNIQUE NOT NULL,
       name TEXT UNIQUE NOT NULL,
       description TEXT NOT NULL,
       position INT NOT NULL,
       created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
       updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
       deleted_AT TIMESTAMP,

       PRIMARY KEY (id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE competitions;
