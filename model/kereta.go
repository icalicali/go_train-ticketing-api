package model

import (
	"gorm.io/gorm"
)

type Kereta struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	Name      string         `json:"name"`
}
