// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: meetings.sql

package board

import (
	"context"
	"time"
)

const createMeeting = `-- name: CreateMeeting :one
INSERT INTO meetings (id, meeting_type, meeting_date)
VALUES ($1, $2, $3) RETURNING id, meeting_type, meeting_date
`

type CreateMeetingParams struct {
	ID          int64     `json:"id"`
	MeetingType string    `json:"meetingType"`
	MeetingDate time.Time `json:"meetingDate"`
}

func (q *Queries) CreateMeeting(ctx context.Context, arg CreateMeetingParams) (Meeting, error) {
	row := q.db.QueryRowContext(ctx, createMeeting, arg.ID, arg.MeetingType, arg.MeetingDate)
	var i Meeting
	err := row.Scan(&i.ID, &i.MeetingType, &i.MeetingDate)
	return i, err
}

const deleteMeeting = `-- name: DeleteMeeting :exec
DELETE FROM meetings WHERE id = $1
`

func (q *Queries) DeleteMeeting(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteMeeting, id)
	return err
}

const getMeetingByID = `-- name: GetMeetingByID :one
SELECT id, meeting_type, meeting_date FROM meetings WHERE id = $1
`

func (q *Queries) GetMeetingByID(ctx context.Context, id int64) (Meeting, error) {
	row := q.db.QueryRowContext(ctx, getMeetingByID, id)
	var i Meeting
	err := row.Scan(&i.ID, &i.MeetingType, &i.MeetingDate)
	return i, err
}

const getMeetingsByDate = `-- name: GetMeetingsByDate :many
SELECT id, meeting_type, meeting_date FROM meetings WHERE meeting_date = $1
`

func (q *Queries) GetMeetingsByDate(ctx context.Context, meetingDate time.Time) ([]Meeting, error) {
	rows, err := q.db.QueryContext(ctx, getMeetingsByDate, meetingDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Meeting
	for rows.Next() {
		var i Meeting
		if err := rows.Scan(&i.ID, &i.MeetingType, &i.MeetingDate); err != nil {
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

const updateMeeting = `-- name: UpdateMeeting :exec
UPDATE meetings SET meeting_type = $2, meeting_date = $3 WHERE id = $1
`

type UpdateMeetingParams struct {
	ID          int64     `json:"id"`
	MeetingType string    `json:"meetingType"`
	MeetingDate time.Time `json:"meetingDate"`
}

func (q *Queries) UpdateMeeting(ctx context.Context, arg UpdateMeetingParams) error {
	_, err := q.db.ExecContext(ctx, updateMeeting, arg.ID, arg.MeetingType, arg.MeetingDate)
	return err
}
