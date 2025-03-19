-- name: CreateAssignment :one
INSERT INTO assignments (id, meeting_id, subsession_id, user_id, assignment_type)
VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: GetAssignmentsByMeeting :many
SELECT * FROM assignments WHERE meeting_id = $1;

-- name: GetAssignmentsBySubsession :many
SELECT * FROM assignments WHERE subsession_id = $1;

-- name: GetAssignmentsPaginated :many
SELECT * FROM assignments
ORDER BY id DESC
LIMIT $1 OFFSET $2;


-- name: UpdateAssignment :exec
UPDATE assignments SET user_id = $2, assignment_type = $3 WHERE id = $1;

-- name: DeleteAssignment :exec
DELETE FROM assignments WHERE id = $1;
