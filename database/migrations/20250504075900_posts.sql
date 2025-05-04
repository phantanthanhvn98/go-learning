-- +migrate Up

CREATE TABLE go_learning.posts (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    body TEXT,
    user_id INTEGER NOT NULL,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES go_learning.users (id) ON DELETE CASCADE
);

-- +migrate Down
DROP TABLE go_learning.posts;