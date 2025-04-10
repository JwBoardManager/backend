// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: assignments.sql

package board

import (
	"context"
	"database/sql"
)

const createAssignment = `-- name: CreateAssignment :one
INSERT INTO assignments (id, meeting_id, subsession_id, user_id, assignment_type)
VALUES ($1, $2, $3, $4, $5) RETURNING id, meeting_id, subsession_id, user_id, assignment_type
`

type CreateAssignmentParams struct {
	ID             int64         `json:"id"`
	MeetingID      sql.NullInt64 `json:"meetingId"`
	SubsessionID   sql.NullInt64 `json:"subsessionId"`
	UserID         int64         `json:"userId"`
	AssignmentType string        `json:"assignmentType"`
}

func (q *Queries) CreateAssignment(ctx context.Context, arg CreateAssignmentParams) (Assignment, error) {
	row := q.db.QueryRowContext(ctx, createAssignment,
		arg.ID,
		arg.MeetingID,
		arg.SubsessionID,
		arg.UserID,
		arg.AssignmentType,
	)
	var i Assignment
	err := row.Scan(
		&i.ID,
		&i.MeetingID,
		&i.SubsessionID,
		&i.UserID,
		&i.AssignmentType,
	)
	return i, err
}

const deleteAssignment = `-- name: DeleteAssignment :exec
DELETE FROM assignments WHERE id = $1
`

func (q *Queries) DeleteAssignment(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteAssignment, id)
	return err
}

const getAssignmentByID = `-- name: GetAssignmentByID :one
SELECT id, meeting_id, subsession_id, user_id, assignment_type FROM assignments WHERE id = $1
`

func (q *Queries) GetAssignmentByID(ctx context.Context, id int64) (Assignment, error) {
	row := q.db.QueryRowContext(ctx, getAssignmentByID, id)
	var i Assignment
	err := row.Scan(
		&i.ID,
		&i.MeetingID,
		&i.SubsessionID,
		&i.UserID,
		&i.AssignmentType,
	)
	return i, err
}

const getAssignmentByUserID = `-- name: GetAssignmentByUserID :many
SELECT id, meeting_id, subsession_id, user_id, assignment_type FROM assignments WHERE user_id = $1
`

func (q *Queries) GetAssignmentByUserID(ctx context.Context, userID int64) ([]Assignment, error) {
	rows, err := q.db.QueryContext(ctx, getAssignmentByUserID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Assignment
	for rows.Next() {
		var i Assignment
		if err := rows.Scan(
			&i.ID,
			&i.MeetingID,
			&i.SubsessionID,
			&i.UserID,
			&i.AssignmentType,
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

const getAssignmentsByMeeting = `-- name: GetAssignmentsByMeeting :many
SELECT id, meeting_id, subsession_id, user_id, assignment_type FROM assignments WHERE meeting_id = $1
`

func (q *Queries) GetAssignmentsByMeeting(ctx context.Context, meetingID sql.NullInt64) ([]Assignment, error) {
	rows, err := q.db.QueryContext(ctx, getAssignmentsByMeeting, meetingID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Assignment
	for rows.Next() {
		var i Assignment
		if err := rows.Scan(
			&i.ID,
			&i.MeetingID,
			&i.SubsessionID,
			&i.UserID,
			&i.AssignmentType,
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

const getAssignmentsBySubsession = `-- name: GetAssignmentsBySubsession :many
SELECT id, meeting_id, subsession_id, user_id, assignment_type FROM assignments WHERE subsession_id = $1
`

func (q *Queries) GetAssignmentsBySubsession(ctx context.Context, subsessionID sql.NullInt64) ([]Assignment, error) {
	rows, err := q.db.QueryContext(ctx, getAssignmentsBySubsession, subsessionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Assignment
	for rows.Next() {
		var i Assignment
		if err := rows.Scan(
			&i.ID,
			&i.MeetingID,
			&i.SubsessionID,
			&i.UserID,
			&i.AssignmentType,
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

const getAssignmentsPaginated = `-- name: GetAssignmentsPaginated :many
SELECT id, meeting_id, subsession_id, user_id, assignment_type FROM assignments
ORDER BY id DESC
LIMIT $1 OFFSET $2
`

type GetAssignmentsPaginatedParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetAssignmentsPaginated(ctx context.Context, arg GetAssignmentsPaginatedParams) ([]Assignment, error) {
	rows, err := q.db.QueryContext(ctx, getAssignmentsPaginated, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Assignment
	for rows.Next() {
		var i Assignment
		if err := rows.Scan(
			&i.ID,
			&i.MeetingID,
			&i.SubsessionID,
			&i.UserID,
			&i.AssignmentType,
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

const getAssignmentsPaginatedCount = `-- name: GetAssignmentsPaginatedCount :one
SELECT COUNT(*) FROM assignments
`

func (q *Queries) GetAssignmentsPaginatedCount(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, getAssignmentsPaginatedCount)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const updateAssignment = `-- name: UpdateAssignment :exec
UPDATE assignments SET user_id = $2, assignment_type = $3 WHERE id = $1
`

type UpdateAssignmentParams struct {
	ID             int64  `json:"id"`
	UserID         int64  `json:"userId"`
	AssignmentType string `json:"assignmentType"`
}

func (q *Queries) UpdateAssignment(ctx context.Context, arg UpdateAssignmentParams) error {
	_, err := q.db.ExecContext(ctx, updateAssignment, arg.ID, arg.UserID, arg.AssignmentType)
	return err
}
