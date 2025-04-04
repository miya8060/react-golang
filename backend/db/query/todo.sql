-- name: GetTodo :one
SELECT * FROM todos
WHERE id = $1 LIMIT 1;

-- name: ListTodos :many
SELECT * FROM todos
ORDER BY id;

-- name: CreateTodo :one
INSERT INTO todos (
    title
) VALUES (
    $1
)
RETURNING *;

-- name: DeleteTodo :exec
DELETE FROM todos
WHERE id = $1;

-- name: UpdateTodo :one
UPDATE todos
SET title = $2
WHERE id = $1
RETURNING *;