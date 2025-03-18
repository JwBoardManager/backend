// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package cleaning

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"
)

type AssignmentTypeEnum string

const (
	AssignmentTypeEnumReading                         AssignmentTypeEnum = "Reading"
	AssignmentTypeEnumDemonstrationStudent            AssignmentTypeEnum = "Demonstration Student"
	AssignmentTypeEnumDemonstrationAssistant          AssignmentTypeEnum = "Demonstration Assistant"
	AssignmentTypeEnumDiscourse                       AssignmentTypeEnum = "Discourse"
	AssignmentTypeEnumPrayer                          AssignmentTypeEnum = "Prayer"
	AssignmentTypeEnumChairman                        AssignmentTypeEnum = "Chairman"
	AssignmentTypeEnumWatchtowerReader                AssignmentTypeEnum = "Watchtower Reader"
	AssignmentTypeEnumWatchtowerConductor             AssignmentTypeEnum = "Watchtower Conductor"
	AssignmentTypeEnumCongregationBibleStudyConductor AssignmentTypeEnum = "Congregation Bible Study Conductor"
	AssignmentTypeEnumCongregationBibleStudyReader    AssignmentTypeEnum = "Congregation Bible Study Reader"
	AssignmentTypeEnumAttendants                      AssignmentTypeEnum = "Attendants"
	AssignmentTypeEnumSoundOperator                   AssignmentTypeEnum = "Sound Operator"
	AssignmentTypeEnumMicrophoneOperator              AssignmentTypeEnum = "Microphone Operator"
	AssignmentTypeEnumVideoOperator                   AssignmentTypeEnum = "Video Operator"
	AssignmentTypeEnumPlatformAssistant               AssignmentTypeEnum = "Platform Assistant"
	AssignmentTypeEnumFieldServiceConductor           AssignmentTypeEnum = "Field Service Conductor"
	AssignmentTypeEnumPublicTalkSpeaker               AssignmentTypeEnum = "Public Talk Speaker"
)

func (e *AssignmentTypeEnum) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = AssignmentTypeEnum(s)
	case string:
		*e = AssignmentTypeEnum(s)
	default:
		return fmt.Errorf("unsupported scan type for AssignmentTypeEnum: %T", src)
	}
	return nil
}

type NullAssignmentTypeEnum struct {
	AssignmentTypeEnum AssignmentTypeEnum `json:"assignmentTypeEnum"`
	Valid              bool               `json:"valid"` // Valid is true if AssignmentTypeEnum is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullAssignmentTypeEnum) Scan(value interface{}) error {
	if value == nil {
		ns.AssignmentTypeEnum, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.AssignmentTypeEnum.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullAssignmentTypeEnum) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.AssignmentTypeEnum), nil
}

type CleaningTypeEnum string

const (
	CleaningTypeEnumBeforeMeeting CleaningTypeEnum = "Before Meeting"
	CleaningTypeEnumAfterMeeting  CleaningTypeEnum = "After Meeting"
)

func (e *CleaningTypeEnum) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = CleaningTypeEnum(s)
	case string:
		*e = CleaningTypeEnum(s)
	default:
		return fmt.Errorf("unsupported scan type for CleaningTypeEnum: %T", src)
	}
	return nil
}

type NullCleaningTypeEnum struct {
	CleaningTypeEnum CleaningTypeEnum `json:"cleaningTypeEnum"`
	Valid            bool             `json:"valid"` // Valid is true if CleaningTypeEnum is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullCleaningTypeEnum) Scan(value interface{}) error {
	if value == nil {
		ns.CleaningTypeEnum, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.CleaningTypeEnum.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullCleaningTypeEnum) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.CleaningTypeEnum), nil
}

type MeetingTypeEnum string

const (
	MeetingTypeEnumMidweekMeeting MeetingTypeEnum = "Midweek Meeting"
	MeetingTypeEnumWeekendMeeting MeetingTypeEnum = "Weekend Meeting"
)

func (e *MeetingTypeEnum) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = MeetingTypeEnum(s)
	case string:
		*e = MeetingTypeEnum(s)
	default:
		return fmt.Errorf("unsupported scan type for MeetingTypeEnum: %T", src)
	}
	return nil
}

type NullMeetingTypeEnum struct {
	MeetingTypeEnum MeetingTypeEnum `json:"meetingTypeEnum"`
	Valid           bool            `json:"valid"` // Valid is true if MeetingTypeEnum is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullMeetingTypeEnum) Scan(value interface{}) error {
	if value == nil {
		ns.MeetingTypeEnum, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.MeetingTypeEnum.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullMeetingTypeEnum) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.MeetingTypeEnum), nil
}

type VisitCategoryEnum string

const (
	VisitCategoryEnumCampaign VisitCategoryEnum = "Campaign"
	VisitCategoryEnumNormal   VisitCategoryEnum = "Normal"
)

func (e *VisitCategoryEnum) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = VisitCategoryEnum(s)
	case string:
		*e = VisitCategoryEnum(s)
	default:
		return fmt.Errorf("unsupported scan type for VisitCategoryEnum: %T", src)
	}
	return nil
}

type NullVisitCategoryEnum struct {
	VisitCategoryEnum VisitCategoryEnum `json:"visitCategoryEnum"`
	Valid             bool              `json:"valid"` // Valid is true if VisitCategoryEnum is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullVisitCategoryEnum) Scan(value interface{}) error {
	if value == nil {
		ns.VisitCategoryEnum, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.VisitCategoryEnum.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullVisitCategoryEnum) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.VisitCategoryEnum), nil
}

type VisitTypeEnum string

const (
	VisitTypeEnumNobodyHome VisitTypeEnum = "Nobody Home"
	VisitTypeEnumBusy       VisitTypeEnum = "Busy"
	VisitTypeEnumVisited    VisitTypeEnum = "Visited"
	VisitTypeEnumLetter     VisitTypeEnum = "Letter"
)

func (e *VisitTypeEnum) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = VisitTypeEnum(s)
	case string:
		*e = VisitTypeEnum(s)
	default:
		return fmt.Errorf("unsupported scan type for VisitTypeEnum: %T", src)
	}
	return nil
}

type NullVisitTypeEnum struct {
	VisitTypeEnum VisitTypeEnum `json:"visitTypeEnum"`
	Valid         bool          `json:"valid"` // Valid is true if VisitTypeEnum is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullVisitTypeEnum) Scan(value interface{}) error {
	if value == nil {
		ns.VisitTypeEnum, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.VisitTypeEnum.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullVisitTypeEnum) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.VisitTypeEnum), nil
}

type WeekdayEnum string

const (
	WeekdayEnumMonday    WeekdayEnum = "Monday"
	WeekdayEnumTuesday   WeekdayEnum = "Tuesday"
	WeekdayEnumWednesday WeekdayEnum = "Wednesday"
	WeekdayEnumThursday  WeekdayEnum = "Thursday"
	WeekdayEnumFriday    WeekdayEnum = "Friday"
	WeekdayEnumSaturday  WeekdayEnum = "Saturday"
	WeekdayEnumSunday    WeekdayEnum = "Sunday"
)

func (e *WeekdayEnum) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = WeekdayEnum(s)
	case string:
		*e = WeekdayEnum(s)
	default:
		return fmt.Errorf("unsupported scan type for WeekdayEnum: %T", src)
	}
	return nil
}

type NullWeekdayEnum struct {
	WeekdayEnum WeekdayEnum `json:"weekdayEnum"`
	Valid       bool        `json:"valid"` // Valid is true if WeekdayEnum is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullWeekdayEnum) Scan(value interface{}) error {
	if value == nil {
		ns.WeekdayEnum, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.WeekdayEnum.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullWeekdayEnum) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.WeekdayEnum), nil
}

type Assignment struct {
	ID             int64         `json:"id"`
	MeetingID      sql.NullInt64 `json:"meetingId"`
	SubsessionID   sql.NullInt64 `json:"subsessionId"`
	UserID         int64         `json:"userId"`
	AssignmentType string        `json:"assignmentType"`
}

type Cart struct {
	ID          int64          `json:"id"`
	Location    string         `json:"location"`
	Description sql.NullString `json:"description"`
}

type CartAssignment struct {
	ID      int64 `json:"id"`
	ShiftID int64 `json:"shiftId"`
	UserID  int64 `json:"userId"`
}

type CartShift struct {
	ID        int64     `json:"id"`
	CartID    int64     `json:"cartId"`
	ShiftDay  int32     `json:"shiftDay"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}

type CleaningAssignment struct {
	ID           int64  `json:"id"`
	GroupID      int64  `json:"groupId"`
	MeetingID    int64  `json:"meetingId"`
	CleaningType string `json:"cleaningType"`
}

type Group struct {
	ID          int64         `json:"id"`
	Name        string        `json:"name"`
	LeaderID    int64         `json:"leaderId"`
	AssistantID sql.NullInt64 `json:"assistantId"`
}

type HouseVisit struct {
	ID            int64     `json:"id"`
	TerritoryID   int64     `json:"territoryId"`
	HouseNumber   string    `json:"houseNumber"`
	VisitDate     time.Time `json:"visitDate"`
	VisitType     string    `json:"visitType"`
	VisitCategory string    `json:"visitCategory"`
}

type Meeting struct {
	ID          int64     `json:"id"`
	MeetingType string    `json:"meetingType"`
	MeetingDate time.Time `json:"meetingDate"`
}

type Room struct {
	ID        int64  `json:"id"`
	MeetingID int64  `json:"meetingId"`
	RoomName  string `json:"roomName"`
}

type Session struct {
	ID            int64 `json:"id"`
	RoomID        int64 `json:"roomId"`
	SessionTypeID int32 `json:"sessionTypeId"`
}

type SessionType struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

type Subsession struct {
	ID               int64 `json:"id"`
	SessionID        int64 `json:"sessionId"`
	SubsessionTypeID int32 `json:"subsessionTypeId"`
	DurationMinutes  int16 `json:"durationMinutes"`
}

type SubsessionType struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

type Territory struct {
	ID              int64        `json:"id"`
	TerritoryNumber int32        `json:"territoryNumber"`
	Location        string       `json:"location"`
	Shapefile       []byte       `json:"shapefile"`
	CompletedAt     sql.NullTime `json:"completedAt"`
}

type User struct {
	ID    int64          `json:"id"`
	Name  string         `json:"name"`
	Email sql.NullString `json:"email"`
	Role  string         `json:"role"`
}

type UserGroup struct {
	UserID  int64 `json:"userId"`
	GroupID int64 `json:"groupId"`
}
