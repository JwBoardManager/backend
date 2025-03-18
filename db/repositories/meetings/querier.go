// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package meetings

import (
	"context"
	"time"
)

type Querier interface {
	CreateMeeting(ctx context.Context, arg CreateMeetingParams) (Meeting, error)
	DeleteMeeting(ctx context.Context, id int64) error
	GetMeetingByID(ctx context.Context, id int64) (Meeting, error)
	GetMeetingsByDate(ctx context.Context, meetingDate time.Time) ([]Meeting, error)
	UpdateMeeting(ctx context.Context, arg UpdateMeetingParams) error
}

var _ Querier = (*Queries)(nil)
