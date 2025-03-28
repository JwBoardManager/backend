// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package board

import (
	"context"
	"database/sql"
	"time"
)

type Querier interface {
	AssignCleaning(ctx context.Context, arg AssignCleaningParams) (CleaningAssignment, error)
	AssignUserToCartShift(ctx context.Context, arg AssignUserToCartShiftParams) (CartAssignment, error)
	CreateAssignment(ctx context.Context, arg CreateAssignmentParams) (Assignment, error)
	CreateCart(ctx context.Context, arg CreateCartParams) (Cart, error)
	CreateCartShift(ctx context.Context, arg CreateCartShiftParams) (CartShift, error)
	CreateGroup(ctx context.Context, arg CreateGroupParams) (Group, error)
	CreateHouseVisit(ctx context.Context, arg CreateHouseVisitParams) (HouseVisit, error)
	CreateMeeting(ctx context.Context, arg CreateMeetingParams) (Meeting, error)
	CreateRoom(ctx context.Context, arg CreateRoomParams) (Room, error)
	CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error)
	CreateSubsession(ctx context.Context, arg CreateSubsessionParams) (Subsession, error)
	CreateTerritory(ctx context.Context, arg CreateTerritoryParams) (Territory, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteAssignment(ctx context.Context, id int64) error
	DeleteCart(ctx context.Context, id int64) error
	DeleteCartAssignment(ctx context.Context, id int64) error
	DeleteCartShift(ctx context.Context, id int64) error
	DeleteCleaningAssignment(ctx context.Context, id int64) error
	DeleteGroup(ctx context.Context, id int64) error
	DeleteHouseVisit(ctx context.Context, id int64) error
	DeleteMeeting(ctx context.Context, id int64) error
	DeleteRoom(ctx context.Context, id int64) error
	DeleteSession(ctx context.Context, id int64) error
	DeleteSubsession(ctx context.Context, id int64) error
	DeleteTerritory(ctx context.Context, id int64) error
	DeleteUser(ctx context.Context, id int64) error
	GetAllCarts(ctx context.Context) ([]Cart, error)
	GetAllTerritories(ctx context.Context) ([]Territory, error)
	GetAllUsers(ctx context.Context) ([]User, error)
	GetAssignmentByID(ctx context.Context, id int64) (Assignment, error)
	GetAssignmentByUserID(ctx context.Context, userID int64) ([]Assignment, error)
	GetAssignmentsByMeeting(ctx context.Context, meetingID sql.NullInt64) ([]Assignment, error)
	GetAssignmentsBySubsession(ctx context.Context, subsessionID sql.NullInt64) ([]Assignment, error)
	GetAssignmentsPaginated(ctx context.Context, arg GetAssignmentsPaginatedParams) ([]Assignment, error)
	GetAssignmentsPaginatedCount(ctx context.Context) (int64, error)
	GetAvailableCartShifts(ctx context.Context, arg GetAvailableCartShiftsParams) ([]sql.NullString, error)
	GetCartAssignmentsByShift(ctx context.Context, shiftID int64) ([]CartAssignment, error)
	GetCartShiftsByDay(ctx context.Context, shiftDay int32) ([]CartShift, error)
	GetCleaningAssignmentsByMeeting(ctx context.Context, meetingID int64) ([]CleaningAssignment, error)
	GetCompletedTerritories(ctx context.Context) ([]Territory, error)
	GetEnumAssignmentTypes(ctx context.Context) ([]interface{}, error)
	GetEnumCleaningTypes(ctx context.Context) ([]interface{}, error)
	GetEnumMeetingTypes(ctx context.Context) ([]interface{}, error)
	GetGroupByID(ctx context.Context, id int64) (Group, error)
	GetHouseVisitsPaginated(ctx context.Context, arg GetHouseVisitsPaginatedParams) ([]HouseVisit, error)
	GetMeetingByID(ctx context.Context, id int64) (Meeting, error)
	GetMeetingsByDate(ctx context.Context, meetingDate time.Time) ([]Meeting, error)
	GetRoomsByMeeting(ctx context.Context, meetingID int64) ([]Room, error)
	GetSessionsByRoom(ctx context.Context, roomID int64) ([]Session, error)
	GetSubsessionsBySession(ctx context.Context, sessionID int64) ([]Subsession, error)
	GetTerritoriesPaginated(ctx context.Context, arg GetTerritoriesPaginatedParams) ([]Territory, error)
	GetTerritoryByID(ctx context.Context, id int64) (Territory, error)
	GetUserByID(ctx context.Context, id int64) (User, error)
	GetVisitsByDate(ctx context.Context, visitDate time.Time) ([]HouseVisit, error)
	GetVisitsByHouseNumber(ctx context.Context, arg GetVisitsByHouseNumberParams) ([]HouseVisit, error)
	GetVisitsByTerritory(ctx context.Context, territoryID int64) ([]HouseVisit, error)
	MarkTerritoryAsCompleted(ctx context.Context, arg MarkTerritoryAsCompletedParams) error
	ReopenTerritory(ctx context.Context, id int64) error
	UpdateAssignment(ctx context.Context, arg UpdateAssignmentParams) error
	UpdateCart(ctx context.Context, arg UpdateCartParams) error
	UpdateCartShift(ctx context.Context, arg UpdateCartShiftParams) error
	UpdateGroup(ctx context.Context, arg UpdateGroupParams) error
	UpdateHouseVisit(ctx context.Context, arg UpdateHouseVisitParams) error
	UpdateMeeting(ctx context.Context, arg UpdateMeetingParams) error
	UpdateRoom(ctx context.Context, arg UpdateRoomParams) error
	UpdateSession(ctx context.Context, arg UpdateSessionParams) error
	UpdateSubsession(ctx context.Context, arg UpdateSubsessionParams) error
	UpdateUser(ctx context.Context, arg UpdateUserParams) error
}

var _ Querier = (*Queries)(nil)
