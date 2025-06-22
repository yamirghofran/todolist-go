-- +goose Up
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  oauth_provider VARCHAR(50),
  oauth_id TEXT,
  name VARCHAR(255),
  email VARCHAR(255) NOT NULL,
  hashed_password CHAR(60),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE users ADD CONSTRAINT unique_user_email UNIQUE (email);

ALTER TABLE todos
ADD COLUMN user_id INT;

ALTER TABLE todos ADD CONSTRAINT fk_todos_user FOREIGN KEY (user_id) REFERENCES users (id);

-- +goose Down
ALTER TABLE todos
DROP CONSTRAINT fk_todos_user;

ALTER TABLE todos
DROP user_id;

ALTER TABLE users
DROP CONSTRAINT unique_user_email;

DROP TABLE users;
