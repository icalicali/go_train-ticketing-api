package method

import (
	"go_mini-project/database"
	"go_mini-project/model"
)

//method1

type KeretaMethod interface {
	GetAllTrain() []model.Kereta
	GetTrainByID(id string) model.Kereta
	CreateTrain(input model.KeretaInput) model.Kereta
	UpdateTrain(id string, input model.KeretaInput) model.Kereta
	DeleteTrain(id string) bool
}

type KeretaMethodCRUD struct{}

func (n *KeretaMethodCRUD) GetAllTrain() []model.Kereta {
	var train []model.Kereta

	database.DB.Find(&train)

	return train
}

func (c *KeretaMethodCRUD) GetTrainByID(id string) model.Kereta {
	var train model.Kereta

	database.DB.First(&train, "id = ?", id)

	return train
}

func (n *KeretaMethodCRUD) CreateTrain(input model.KeretaInput) model.Kereta {
	var newTrain model.Kereta = model.Kereta{
		Name: input.Name,
	}

	var createdTrain model.Kereta = model.Kereta{}

	result := database.DB.Create(&newTrain)

	result.Last(&createdTrain)

	return createdTrain
}

func (n *KeretaMethodCRUD) UpdateTrain(id string, input model.KeretaInput) model.Kereta {
	var train model.Kereta = n.GetTrainByID(id)

	train.Name = input.Name

	database.DB.Save(&train)

	return train
}

func (n *KeretaMethodCRUD) DeleteTrain(id string) bool {
	var train model.Kereta = n.GetTrainByID(id)

	result := database.DB.Delete(&train)

	if result.RowsAffected == 0 {
		return false
	}

	return true
}
