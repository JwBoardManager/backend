-- name: CreateCart :one
INSERT INTO carts (id, location, description)
VALUES ($1, $2, $3) RETURNING *;

-- name: GetAllCarts :many
SELECT * FROM carts;

-- name: UpdateCart :exec
UPDATE carts SET location = $2, description = $3 WHERE id = $1;

-- name: DeleteCart :exec
DELETE FROM carts WHERE id = $1;
