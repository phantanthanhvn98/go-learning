-- +migrate Up
CREATE TABLE go_learning.users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL
);

-- +migrate Down
DROP TABLE go_learning.users;