CREATE TABLE project
(
    id          BIGSERIAL PRIMARY KEY,
    title       TEXT NOT NULL,
    description text NOT NULL,
    github_url  text,
    project_url text,
    created_at  TIMESTAMP DEFAULT NOW()
);