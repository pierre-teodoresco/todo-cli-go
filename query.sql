-- name: FindAllTasks :many
SELECT * FROM tasks ORDER BY created_at;

-- name: FindAllUnfinishedTasks :many
SELECT * FROM tasks WHERE completed = 0 ORDER BY created_at;

-- name: NewTask :one
INSERT INTO tasks (title)
VALUES (?)
RETURNING *;

-- name: CompleteTask :one
UPDATE tasks
SET completed = 1
WHERE id = ?
RETURNING *;

-- name: ReverseCompleteTask :one
UPDATE tasks
SET completed = 0
WHERE id = ?
RETURNING *;

-- name: DeleteTask :exec
DELETE FROM tasks
WHERE id = ?;
