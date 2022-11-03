package method

import (
	"go_mini-project/database"
	"go_mini-project/model"
)

type TiketMethodCRUD struct{}

func (n *TiketMethodCRUD) GetAll() []model.Tiket {
	var tiket []model.Tiket

	database.DB.Find(&tiket)

	return tiket
}

func (c *TiketMethodCRUD) GetByID(id string) model.Tiket {
	var tiket model.Tiket

	database.DB.First(&tiket, "id = ?", id)

	return tiket
}

func (n *TiketMethodCRUD) Create(input model.TiketInput) model.Tiket {
	var newTiket model.Tiket = model.Tiket{
		Title:       input.Title,
		Description: input.Description,
	}

	var createdTiket model.Tiket = model.Tiket{}

	result := database.DB.Create(&newTiket)

	result.Last(&createdTiket)

	return createdTiket
}

func (n *TiketMethodCRUD) Update(id string, input model.TiketInput) model.Tiket {
	var tiket model.Tiket = n.GetByID(id)

	tiket.Title = input.Title
	tiket.Description = input.Description

	database.DB.Save(&tiket)

	return tiket
}

func (n *TiketMethodCRUD) Delete(id string) bool {
	var tiket model.Tiket = n.GetByID(id)

	result := database.DB.Delete(&tiket)

	if result.RowsAffected == 0 {
		return false
	}

	return true
}
