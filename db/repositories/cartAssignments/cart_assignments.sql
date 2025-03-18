-- name: AssignUserToCartShift :one
INSERT INTO cart_assignments (id, shift_id, user_id)
VALUES ($1, $2, $3) RETURNING *;

-- name: GetCartAssignmentsByShift :many
SELECT * FROM cart_assignments WHERE shift_id = $1;

-- name: DeleteCartAssignment :exec
DELETE FROM cart_assignments WHERE id = $1;
