package model

import "github.com/go-playground/validator/v10"

type TiketInput struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	KeretaID    uint   `json:"kereta_id" validate:"required"`
}

func (input *TiketInput) Validate() error {
	validate := validator.New()

	err := validate.Struct(input)

	return err
}
