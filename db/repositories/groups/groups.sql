-- name: CreateGroup :one
INSERT INTO groups (id, name, leader_id, assistant_id)
VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetGroupByID :one
SELECT * FROM groups WHERE id = $1;

-- name: UpdateGroup :exec
UPDATE groups SET name = $2, leader_id = $3, assistant_id = $4 WHERE id = $1;

-- name: DeleteGroup :exec
DELETE FROM groups WHERE id = $1;
