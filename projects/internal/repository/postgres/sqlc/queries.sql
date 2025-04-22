-- name: GetProject :one
SELECT id, title, description, github_url, project_url, created_at
FROM project
WHERE id = $1 LIMIT 1;

-- name: ListProjects :many
SELECT id, title, description, github_url, project_url, created_at
FROM project LIMIT $1
OFFSET $2;

-- name: CreateProject :one
INSERT INTO project (title, description, github_url, project_url)
VALUES ($1, $2, $3, $4) RETURNING id;

-- name: DeleteProject :exec
DELETE
FROM project
WHERE id = $1;

-- name: UpdateProject :exec
UPDATE project
SET title = $1,
    description = $2,
    github_url = $3,
    project_url = $4
WHERE id = $5;
