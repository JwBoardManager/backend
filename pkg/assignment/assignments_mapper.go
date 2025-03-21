package assignment

import (
	board "backend/db/sqlc/out"
	"backend/pkg/models"
)

// AssignmentResponse representa a resposta de uma atribuição
type AssignmentResponse struct {
	ID             int64  `json:"id"`             // ID da atribuição
	MeetingID      int64  `json:"meetingId"`      // ID da reunião
	SubsessionID   int64  `json:"subsessionId"`   // ID da subsessão
	UserID         int64  `json:"userId"`         // ID do usuário
	AssignmentType string `json:"assignmentType"` // Tipo de atribuição
}

// Representa os dados de paginação
type Pagination struct {
	Page       int32 `json:"page"`       // Página atual
	PageSize   int32 `json:"pageSize"`   // Itens por página
	TotalItems int32 `json:"totalItems"` // Total de itens
	TotalPages int32 `json:"totalPages"` // Total de páginas
	HasNext    bool  `json:"hasNext"`    // Existe próxima página
	HasPrev    bool  `json:"hasPrev"`    // Existe página anterior
}

// Resposta paginada com assignments
type AssignmentResponsePaginated struct {
	Data       []AssignmentResponse `json:"data"`
	Pagination Pagination           `json:"pagination"`
}

// Converte `out.Assignment` para `models.AssignmentResponse`
func ToAssignmentResponse(a board.Assignment) AssignmentResponse {
	return AssignmentResponse{
		ID:             a.ID,
		MeetingID:      models.FromNullInt64(&a.MeetingID),
		SubsessionID:   models.FromNullInt64(&a.SubsessionID),
		UserID:         a.UserID,
		AssignmentType: a.AssignmentType,
	}
}

// Converte lista de board.Assignment para []AssignmentResponse
func ToAssignmentResponseList(assignments []board.Assignment) []AssignmentResponse {
	responses := make([]AssignmentResponse, 0, len(assignments))
	for _, a := range assignments {
		responses = append(responses, AssignmentResponse{
			ID:             a.ID,
			MeetingID:      models.FromNullInt64(&a.MeetingID),
			SubsessionID:   models.FromNullInt64(&a.SubsessionID),
			UserID:         a.UserID,
			AssignmentType: a.AssignmentType,
		})
	}
	return responses
}

func ToCreateAssignmentParams(req CreateAssignmentDTO) board.CreateAssignmentParams {
	return board.CreateAssignmentParams{
		MeetingID:      models.ToNullInt64(req.MeetingID),
		SubsessionID:   models.ToNullInt64(req.SubsessionID),
		UserID:         req.UserID,
		AssignmentType: req.AssignmentType,
	}
}
