package model

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Tiket struct {
	ID          uint           `json:"id" gorm:"primarykey"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Kereta      Kereta         `json:"kereta"`
	KeretaID    uint           `json:"kereta_id"`
}

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

// ---------------------------------------------------- //

type Kereta struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type KeretaInput struct {
	Name string `json:"name"`
}

func (input *KeretaInput) Validate() error {
	validate := validator.New()

	err := validate.Struct(input)

	return err
}

// ---------------------------------------------------- //

type Pemesan struct {
	ID       uint   `json:"id" gorm:"primarykey"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// ---------------------------------------------------- //

type Register struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (input *Register) Validate() error {
	validate := validator.New()

	err := validate.Struct(input)

	return err
}

// ---------------------------------------------------- //
