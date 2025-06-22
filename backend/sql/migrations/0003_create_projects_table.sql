-- +goose Up
CREATE TABLE projects (
  id SERIAL PRIMARY KEY,
  user_id INT,
  title VARCHAR(255) NOT NULL,
  is_completed BOOLEAN DEFAULT FALSE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE projects ADD CONSTRAINT fk_project_user FOREIGN KEY (user_id) REFERENCES users (id);

ALTER TABLE todos
ADD COLUMN project_id INT;

ALTER TABLE todos ADD CONSTRAINT fk_todos_project FOREIGN KEY (project_id) REFERENCES projects (id);

-- +goose Down
ALTER TABLE todos
DROP CONSTRAINT fk_todos_project;

ALTER TABLE todos
DROP COLUMN project_id;

ALTER TABLE projects
DROP CONSTRAINT fk_project_user;

DROP TABLE projects;
