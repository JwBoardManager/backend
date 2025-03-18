-- name: CreateSession :one
INSERT INTO sessions (id, room_id, session_type_id)
VALUES ($1, $2, $3) RETURNING *;

-- name: GetSessionsByRoom :many
SELECT * FROM sessions WHERE room_id = $1;

-- name: UpdateSession :exec
UPDATE sessions SET session_type_id = $2 WHERE id = $1;

-- name: DeleteSession :exec
DELETE FROM sessions WHERE id = $1;
