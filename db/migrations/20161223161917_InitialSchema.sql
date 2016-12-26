-- +goose Up

CREATE TABLE users (
       id SERIAL UNIQUE NOT NULL,
       email TEXT UNIQUE NOT NULL,
       name TEXT UNIQUE NOT NULL,
       password_hash TEXT NOT NULL,
       created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
       updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
       deleted_at timestamp,

       PRIMARY KEY (id)
);

-- +goose Down

DROP TABLE users;
