// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: cart_assignments.sql

package board

import (
	"context"
)

const assignUserToCartShift = `-- name: AssignUserToCartShift :one
INSERT INTO cart_assignments (id, shift_id, user_id)
VALUES ($1, $2, $3) RETURNING id, shift_id, user_id
`

type AssignUserToCartShiftParams struct {
	ID      int64 `json:"id"`
	ShiftID int64 `json:"shiftId"`
	UserID  int64 `json:"userId"`
}

func (q *Queries) AssignUserToCartShift(ctx context.Context, arg AssignUserToCartShiftParams) (CartAssignment, error) {
	row := q.db.QueryRowContext(ctx, assignUserToCartShift, arg.ID, arg.ShiftID, arg.UserID)
	var i CartAssignment
	err := row.Scan(&i.ID, &i.ShiftID, &i.UserID)
	return i, err
}

const deleteCartAssignment = `-- name: DeleteCartAssignment :exec
DELETE FROM cart_assignments WHERE id = $1
`

func (q *Queries) DeleteCartAssignment(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteCartAssignment, id)
	return err
}

const getCartAssignmentsByShift = `-- name: GetCartAssignmentsByShift :many
SELECT id, shift_id, user_id FROM cart_assignments WHERE shift_id = $1
`

func (q *Queries) GetCartAssignmentsByShift(ctx context.Context, shiftID int64) ([]CartAssignment, error) {
	rows, err := q.db.QueryContext(ctx, getCartAssignmentsByShift, shiftID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []CartAssignment
	for rows.Next() {
		var i CartAssignment
		if err := rows.Scan(&i.ID, &i.ShiftID, &i.UserID); err != nil {
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
