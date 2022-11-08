package model

import "github.com/go-playground/validator/v10"

type Register struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (input *Register) Validate() error {
	validate := validator.New()

	err := validate.Struct(input)

	return err
}
