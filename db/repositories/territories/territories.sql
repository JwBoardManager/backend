-- name: CreateTerritory :one
INSERT INTO territories (id, territory_number, location, shapefile, completed_at)
VALUES ($1, $2, $3, $4, NULL) RETURNING *;

-- name: GetTerritoryByID :one
SELECT * FROM territories WHERE id = $1;

-- name: GetAllTerritories :many
SELECT * FROM territories;

-- name: GetCompletedTerritories :many
SELECT * FROM territories WHERE completed_at IS NOT NULL;

-- name: GetTerritoriesPaginated :many
SELECT * FROM territories
ORDER BY completed_at NULLS FIRST, id DESC
LIMIT $1 OFFSET $2;


-- name: MarkTerritoryAsCompleted :exec
UPDATE territories SET completed_at = $2 WHERE id = $1;

-- name: ReopenTerritory :exec
UPDATE territories SET completed_at = NULL WHERE id = $1;

-- name: DeleteTerritory :exec
DELETE FROM territories WHERE id = $1;
