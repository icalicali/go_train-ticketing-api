package model

type Kereta struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}
