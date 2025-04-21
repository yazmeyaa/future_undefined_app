-- name: GetTaskByID :one
SELECT id,
       title,
       description,
       user_id,
       task_list_id,
       status,
       created_at
FROM task
WHERE id = $1;

-- name: GetTasksByUserID :many
SELECT id,
       title,
       description,
       user_id,
       task_list_id,
       status,
       created_at
FROM task
WHERE user_id = $1
ORDER BY created_at DESC LIMIT $2
OFFSET $3;

-- name: GetTasksByTaskListID :many
SELECT id,
       title,
       description,
       user_id,
       task_list_id,
       status,
       created_at
FROM task
WHERE task_list_id = $1
ORDER BY created_at DESC LIMIT $2
OFFSET $3;

-- name: DeleteTask :exec
DELETE
FROM task
WHERE id = $1;

-- name: DeleteTaskBatch :exec
DELETE
FROM task
WHERE id = ANY ($1::int[]);

-- name: UpdateTask :exec
UPDATE task
SET title        = $1,
    description  = $2,
    user_id      = $3,
    task_list_id = $4,
    status       = $5
WHERE id = $6;

-- name: GetTaskListsByUserID :many
SELECT id,
       user_id,
       title,
       description,
       created_at
FROM task_list
WHERE user_id = $1
ORDER BY created_at DESC LIMIT $2
OFFSET $3;

-- name: GetTaskListByID :one
SELECT id,
       user_id,
       title,
       description,
       created_at
FROM task_list
WHERE id = $1;

-- name: DeleteTaskListByID :exec
DELETE
FROM task_list
WHERE id = $1;

-- name: CreateTask :one
INSERT
INTO task (user_id, task_list_id, title, description, status)
VALUES ($1, $2, $3, $4, $5) RETURNING id;

-- name: CreateTaskList :one
INSERT
INTO task_list (user_id, title, description)
VALUES ($1, $2, $3) RETURNING id;

-- name: UpdateTaskList :exec
UPDATE task_list
SET user_id = CASE WHEN sqlc.narg('user_id')::int IS NOT NULL THEN sqlc.arg('user_id')::int ELSE user_id END,
    title     = CASE WHEN sqlc.narg('title')::text IS NOT NULL THEN sqlc.arg('title')::text ELSE title
END,
    description     = CASE WHEN sqlc.narg('description')::text IS NOT NULL THEN sqlc.arg('description')::text ELSE description
END
WHERE id = $1;



