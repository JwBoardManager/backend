package users

import (
	board "backend/db/sqlc/out"
	"backend/pkg/models"
)

func ToCretaeUserParams(req CreateUserDTO) board.CreateUserParams {
	return board.CreateUserParams{
		Name:     req.Name,
		Email:    models.ToNullString(&req.Email),
		Password: req.Password,
		Role:     board.UserRoleEnum(req.Role),
	}
}
