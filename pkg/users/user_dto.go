package users

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type CreateUserDTO struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role" validate:"required"`
}

type UpdateUserDTO struct {
	ID       int64  `json:"id" validate:"required"`
	Name     string `json:"name"`
	Email    string `json:"email" validate:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

var validRolesUer = map[string]bool{
	"Elder":               true,
	"Ministerial_Servant": true,
	"auxiliar_pioneer":    true,
	"regular_pioneer":     true,
	"student":             true,
	"publisher":           true,
}

var validate = validator.New()

func (c *CreateUserDTO) Validate() error {
	err := validate.Struct(c)
	if err != nil {
		return err
	}

	if c.Role != "" {
		if _, ok := validRolesUer[c.Role]; !ok {
			return errors.New("invalid role user")
		}
	}

	return nil
}

func (u *UpdateUserDTO) Validate() error {
	err := validate.Struct(u)
	if err != nil {
		return err
	}

	return nil
}
