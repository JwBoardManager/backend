package assignment

import (
	"context"
	"database/sql"
	"fmt"

	board "backend/db/sqlc/out"
)

// AssignmentService encapsula a lógica de negócio para atribuições
type AssignmentService struct {
	queries *board.Queries
}

// NewAssignmentService cria um novo serviço de assignments
func NewAssignmentService(queries *board.Queries) *AssignmentService {
	return &AssignmentService{queries: queries}
}

// Criar uma nova atribuição (assignment)
func (s *AssignmentService) CreateAssignment(ctx context.Context, params board.CreateAssignmentParams) (board.Assignment, error) {
	assignment, err := s.queries.CreateAssignment(ctx, params)
	if err != nil {
		return board.Assignment{}, fmt.Errorf("error creating assignment: %w", err)
	}
	return assignment, nil
}

// Buscar atribuições por reunião
func (s *AssignmentService) GetAssignmentsByMeeting(ctx context.Context, meetingID int64) ([]board.Assignment, error) {
	assignments, err := s.queries.GetAssignmentsByMeeting(ctx, sql.NullInt64{Int64: meetingID, Valid: true})
	if err != nil {
		return nil, fmt.Errorf("error fetching assignments by meeting: %w", err)
	}
	return assignments, nil
}

// Buscar atribuições por subsessão
func (s *AssignmentService) GetAssignmentsBySubsession(ctx context.Context, subsessionID int64) ([]board.Assignment, error) {
	assignments, err := s.queries.GetAssignmentsBySubsession(ctx, sql.NullInt64{Int64: subsessionID, Valid: true})
	if err != nil {
		return nil, fmt.Errorf("error fetching assignments by subsession: %w", err)
	}
	return assignments, nil
}

// Buscar atribuições com paginação
func (s *AssignmentService) GetAssignmentsPaginated(ctx context.Context, limit, offset int32) ([]board.Assignment, int32, error) {
	assignments, err := s.queries.GetAssignmentsPaginated(ctx, board.GetAssignmentsPaginatedParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, 0, fmt.Errorf("error fetching paginated assignments: %w", err)
	}

	total, err := s.queries.GetAssignmentsPaginatedCount(ctx)
	if err != nil {
		return nil, 0, fmt.Errorf("error fetching paginated assignments count: %w", err)
	}

	return assignments, int32(total), nil
}

// Atualizar uma atribuição
func (s *AssignmentService) UpdateAssignment(ctx context.Context, params board.UpdateAssignmentParams) error {
	err := s.queries.UpdateAssignment(ctx, params)
	if err != nil {
		return fmt.Errorf("error updating assignment: %w", err)
	}
	return nil
}

// Deletar uma atribuição
func (s *AssignmentService) DeleteAssignment(ctx context.Context, id int64) error {
	err := s.queries.DeleteAssignment(ctx, id)
	if err != nil {
		return fmt.Errorf("error deleting assignment: %w", err)
	}
	return nil
}

// Buscar atribuições por usuário
func (s *AssignmentService) GetAssignmentsByUserID(ctx context.Context, userID int64) ([]board.Assignment, error) {
	assignments, err := s.queries.GetAssignmentByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("error fetching assignments by user: %w", err)
	}
	return assignments, nil
}
