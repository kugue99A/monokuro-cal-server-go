-- name: GetEvent :one
SELECT id, title, description, start_at, end_at, created_at, updated_at
FROM events
WHERE id = $1;

-- name: ListEvents :many
SELECT id, title, description, start_at, end_at, created_at, updated_at
FROM events
ORDER BY start_at;

-- name: CreateEvent :one
INSERT INTO events (id, title, description, start_at, end_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, title, description, start_at, end_at, created_at, updated_at;

-- name: UpdateEvent :one
UPDATE events
SET title = $2, description = $3, start_at = $4, end_at = $5, updated_at = NOW()
WHERE id = $1
RETURNING id, title, description, start_at, end_at, created_at, updated_at;

-- name: DeleteEvent :exec
DELETE FROM events WHERE id = $1;
