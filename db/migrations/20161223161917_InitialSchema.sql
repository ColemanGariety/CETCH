-- +goose Up

CREATE TABLE users (
       id SERIAL UNIQUE NOT NULL,
       name TEXT NOT NULL,
       password_hash TEXT NOT NULL,
       created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
       updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
       deleted_at timestamp,

       PRIMARY KEY (id)
);

CREATE TABLE sessions (
       key TEXT UNIQUE NOT NULL,
       user_id INTEGER NOT NULL,
       created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
       updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
       deleted_at timestamp,

       PRIMARY KEY (key)
);

-- +goose Down

DROP TABLE users;
DROP TABLE sessions;
