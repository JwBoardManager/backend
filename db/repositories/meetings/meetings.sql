-- name: CreateMeeting :one
INSERT INTO meetings (id, meeting_type, meeting_date)
VALUES ($1, $2, $3) RETURNING *;

-- name: GetMeetingByID :one
SELECT * FROM meetings WHERE id = $1;

-- name: GetMeetingsByDate :many
SELECT * FROM meetings WHERE meeting_date = $1;

-- name: UpdateMeeting :exec
UPDATE meetings SET meeting_type = $2, meeting_date = $3 WHERE id = $1;

-- name: DeleteMeeting :exec
DELETE FROM meetings WHERE id = $1;
