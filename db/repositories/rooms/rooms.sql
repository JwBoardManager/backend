-- name: CreateRoom :one
INSERT INTO rooms (id, meeting_id, room_name)
VALUES ($1, $2, $3) RETURNING *;

-- name: GetRoomsByMeeting :many
SELECT * FROM rooms WHERE meeting_id = $1;

-- name: UpdateRoom :exec
UPDATE rooms SET room_name = $2 WHERE id = $1;

-- name: DeleteRoom :exec
DELETE FROM rooms WHERE id = $1;
