// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package assignments

import (
	"context"
	"database/sql"
)

type Querier interface {
	CreateAssignment(ctx context.Context, arg CreateAssignmentParams) (Assignment, error)
	DeleteAssignment(ctx context.Context, id int64) error
	GetAssignmentsByMeeting(ctx context.Context, meetingID sql.NullInt64) ([]Assignment, error)
	GetAssignmentsBySubsession(ctx context.Context, subsessionID sql.NullInt64) ([]Assignment, error)
	GetAssignmentsPaginated(ctx context.Context, arg GetAssignmentsPaginatedParams) ([]Assignment, error)
	UpdateAssignment(ctx context.Context, arg UpdateAssignmentParams) error
}

var _ Querier = (*Queries)(nil)
