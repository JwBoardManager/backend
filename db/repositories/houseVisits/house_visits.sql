-- name: CreateHouseVisit :one
INSERT INTO house_visits (id, territory_id, house_number, visit_date, visit_type, visit_category)
VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: GetVisitsByTerritory :many
SELECT * FROM house_visits WHERE territory_id = $1;

-- name: GetVisitsByHouseNumber :many
SELECT * FROM house_visits WHERE territory_id = $1 AND house_number = $2;

-- name: GetVisitsByDate :many
SELECT * FROM house_visits WHERE visit_date = $1;

-- name: GetHouseVisitsPaginated :many
SELECT * FROM house_visits
WHERE territory_id = $1
ORDER BY visit_date DESC, house_number ASC
LIMIT $2 OFFSET $3;


-- name: UpdateHouseVisit :exec
UPDATE house_visits SET visit_type = $2, visit_category = $3 WHERE id = $1;

-- name: DeleteHouseVisit :exec
DELETE FROM house_visits WHERE id = $1;
