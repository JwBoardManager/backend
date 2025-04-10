// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: house_visits.sql

package board

import (
	"context"
	"time"
)

const createHouseVisit = `-- name: CreateHouseVisit :one
INSERT INTO house_visits (id, territory_id, house_number, visit_date, visit_type, visit_category)
VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, territory_id, house_number, visit_date, visit_type, visit_category
`

type CreateHouseVisitParams struct {
	ID            int64     `json:"id"`
	TerritoryID   int64     `json:"territoryId"`
	HouseNumber   string    `json:"houseNumber"`
	VisitDate     time.Time `json:"visitDate"`
	VisitType     string    `json:"visitType"`
	VisitCategory string    `json:"visitCategory"`
}

func (q *Queries) CreateHouseVisit(ctx context.Context, arg CreateHouseVisitParams) (HouseVisit, error) {
	row := q.db.QueryRowContext(ctx, createHouseVisit,
		arg.ID,
		arg.TerritoryID,
		arg.HouseNumber,
		arg.VisitDate,
		arg.VisitType,
		arg.VisitCategory,
	)
	var i HouseVisit
	err := row.Scan(
		&i.ID,
		&i.TerritoryID,
		&i.HouseNumber,
		&i.VisitDate,
		&i.VisitType,
		&i.VisitCategory,
	)
	return i, err
}

const deleteHouseVisit = `-- name: DeleteHouseVisit :exec
DELETE FROM house_visits WHERE id = $1
`

func (q *Queries) DeleteHouseVisit(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteHouseVisit, id)
	return err
}

const getHouseVisitsPaginated = `-- name: GetHouseVisitsPaginated :many
SELECT id, territory_id, house_number, visit_date, visit_type, visit_category FROM house_visits
WHERE territory_id = $1
ORDER BY visit_date DESC, house_number ASC
LIMIT $2 OFFSET $3
`

type GetHouseVisitsPaginatedParams struct {
	TerritoryID int64 `json:"territoryId"`
	Limit       int32 `json:"limit"`
	Offset      int32 `json:"offset"`
}

func (q *Queries) GetHouseVisitsPaginated(ctx context.Context, arg GetHouseVisitsPaginatedParams) ([]HouseVisit, error) {
	rows, err := q.db.QueryContext(ctx, getHouseVisitsPaginated, arg.TerritoryID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []HouseVisit
	for rows.Next() {
		var i HouseVisit
		if err := rows.Scan(
			&i.ID,
			&i.TerritoryID,
			&i.HouseNumber,
			&i.VisitDate,
			&i.VisitType,
			&i.VisitCategory,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getVisitsByDate = `-- name: GetVisitsByDate :many
SELECT id, territory_id, house_number, visit_date, visit_type, visit_category FROM house_visits WHERE visit_date = $1
`

func (q *Queries) GetVisitsByDate(ctx context.Context, visitDate time.Time) ([]HouseVisit, error) {
	rows, err := q.db.QueryContext(ctx, getVisitsByDate, visitDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []HouseVisit
	for rows.Next() {
		var i HouseVisit
		if err := rows.Scan(
			&i.ID,
			&i.TerritoryID,
			&i.HouseNumber,
			&i.VisitDate,
			&i.VisitType,
			&i.VisitCategory,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getVisitsByHouseNumber = `-- name: GetVisitsByHouseNumber :many
SELECT id, territory_id, house_number, visit_date, visit_type, visit_category FROM house_visits WHERE territory_id = $1 AND house_number = $2
`

type GetVisitsByHouseNumberParams struct {
	TerritoryID int64  `json:"territoryId"`
	HouseNumber string `json:"houseNumber"`
}

func (q *Queries) GetVisitsByHouseNumber(ctx context.Context, arg GetVisitsByHouseNumberParams) ([]HouseVisit, error) {
	rows, err := q.db.QueryContext(ctx, getVisitsByHouseNumber, arg.TerritoryID, arg.HouseNumber)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []HouseVisit
	for rows.Next() {
		var i HouseVisit
		if err := rows.Scan(
			&i.ID,
			&i.TerritoryID,
			&i.HouseNumber,
			&i.VisitDate,
			&i.VisitType,
			&i.VisitCategory,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getVisitsByTerritory = `-- name: GetVisitsByTerritory :many
SELECT id, territory_id, house_number, visit_date, visit_type, visit_category FROM house_visits WHERE territory_id = $1
`

func (q *Queries) GetVisitsByTerritory(ctx context.Context, territoryID int64) ([]HouseVisit, error) {
	rows, err := q.db.QueryContext(ctx, getVisitsByTerritory, territoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []HouseVisit
	for rows.Next() {
		var i HouseVisit
		if err := rows.Scan(
			&i.ID,
			&i.TerritoryID,
			&i.HouseNumber,
			&i.VisitDate,
			&i.VisitType,
			&i.VisitCategory,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateHouseVisit = `-- name: UpdateHouseVisit :exec
UPDATE house_visits SET visit_type = $2, visit_category = $3 WHERE id = $1
`

type UpdateHouseVisitParams struct {
	ID            int64  `json:"id"`
	VisitType     string `json:"visitType"`
	VisitCategory string `json:"visitCategory"`
}

func (q *Queries) UpdateHouseVisit(ctx context.Context, arg UpdateHouseVisitParams) error {
	_, err := q.db.ExecContext(ctx, updateHouseVisit, arg.ID, arg.VisitType, arg.VisitCategory)
	return err
}
