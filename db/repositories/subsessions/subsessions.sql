-- name: CreateSubsession :one
INSERT INTO subsessions (id, session_id, subsession_type_id, duration_minutes)
VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetSubsessionsBySession :many
SELECT * FROM subsessions WHERE session_id = $1;

-- name: UpdateSubsession :exec
UPDATE subsessions SET subsession_type_id = $2, duration_minutes = $3 WHERE id = $1;

-- name: DeleteSubsession :exec
DELETE FROM subsessions WHERE id = $1;
