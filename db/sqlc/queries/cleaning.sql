-- name: AssignCleaning :one
INSERT INTO cleaning_assignments (id, group_id, meeting_id, cleaning_type)
VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetCleaningAssignmentsByMeeting :many
SELECT * FROM cleaning_assignments WHERE meeting_id = $1;

-- name: DeleteCleaningAssignment :exec
DELETE FROM cleaning_assignments WHERE id = $1;
