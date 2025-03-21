package assignment

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

// Lista dos valores permitidos para o `assignment_type_enum`
var validAssignmentTypes = map[string]bool{
	"Reading":                            true,
	"Demonstration Student":              true,
	"Demonstration Assistant":            true,
	"Discourse":                          true,
	"Prayer":                             true,
	"Chairman":                           true,
	"Watchtower Reader":                  true,
	"Watchtower Conductor":               true,
	"Congregation Bible Study Conductor": true,
	"Congregation Bible Study Reader":    true,
	"Attendants":                         true,
	"Sound Operator":                     true,
	"Microphone Operator":                true,
	"Video Operator":                     true,
	"Platform Assistant":                 true,
	"Field Service Conductor":            true,
	"Public Talk Speaker":                true,
}

// DTO para criar um assignment
type CreateAssignmentDTO struct {
	MeetingID      *int64 `json:"meetingId" validate:"required"`
	SubsessionID   *int64 `json:"subsessionId"`
	UserID         int64  `json:"userId" validate:"required"`
	AssignmentType string `json:"assignmentType" validate:"required"`
}

// DTO para atualizar um assignment
type UpdateAssignmentDTO struct {
	ID             int64  `json:"id" validate:"required"`
	UserID         int64  `json:"userId"`
	AssignmentType string `json:"assignmentType"`
}

// DTO para paginação de assignments
type GetAssignmentsPaginatedDTO struct {
	Limit  int `form:"limit" binding:"required,min=1"`
	Offset int `form:"offset" binding:"min=0"`
}

// Inicializa o validador
var validate = validator.New()

// 🚀 Validação personalizada para garantir que meetingId OU subsessionId está preenchido
func (c *CreateAssignmentDTO) Validate() error {
	err := validate.Struct(c)
	if err != nil {
		return err
	}

	if (c.MeetingID != nil && c.SubsessionID != nil) || (c.MeetingID == nil && c.SubsessionID == nil) {
		return errors.New("either meetingId or subsessionId must be provided, but not both")
	}

	if !validAssignmentTypes[c.AssignmentType] {
		return errors.New("invalid assignmentType. Allowed values: " + validAssignmentTypesList())
	}

	return nil
}

// Validação para o UpdateAssignmentDTO
func (u *UpdateAssignmentDTO) Validate() error {
	err := validate.Struct(u)
	if err != nil {
		return err
	}

	// Validação do enum assignment_type_enum
	if u.AssignmentType != "" && !validAssignmentTypes[u.AssignmentType] {
		return errors.New("assignmentType inválido. Valores permitidos: " + validAssignmentTypesList())
	}

	return nil
}

// Função auxiliar para retornar os valores permitidos como string
func validAssignmentTypesList() string {
	keys := make([]string, 0, len(validAssignmentTypes))
	for key := range validAssignmentTypes {
		keys = append(keys, key)
	}
	return "\"" + keys[0] + "\", \"" + keys[1] + "\", ..., \"" + keys[len(keys)-1] + "\""
}
