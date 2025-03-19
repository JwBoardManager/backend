-- name: CreateUser :one
INSERT INTO users (id, name, email, role)
VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;

-- name: GetAllUsers :many
SELECT * FROM users;

-- name: UpdateUser :exec
UPDATE users SET name = $2, email = $3, role = $4 WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;
