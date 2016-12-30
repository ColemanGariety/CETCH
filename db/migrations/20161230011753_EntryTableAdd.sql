-- +goose Up

CREATE TABLE entries (
       id SERIAL UNIQUE NOT NULL,
       user_id int,
       competition_id int,
       created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
       updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
       deleted_at timestamp,

       PRIMARY KEY (id)
);

-- +goose Down

DROP TABLE entries;
