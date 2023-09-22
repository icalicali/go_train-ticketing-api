package service

import (
	"go_mini-project/method"
	"go_mini-project/model"
)

//trainSvc

type KeretaService struct {
	Method method.KeretaMethod
}

func NewTrain() KeretaService {
	return KeretaService{
		Method: &method.KeretaMethodCRUD{},
	}
}

func (n *KeretaService) GetAll() []model.Kereta {
	return n.Method.GetAllTrain()
}

func (n *KeretaService) GetByID(id string) model.Kereta {
	return n.Method.GetTrainByID(id)
}

func (n *KeretaService) Create(input model.KeretaInput) model.Kereta {
	return n.Method.CreateTrain(input)
}

func (n *KeretaService) Update(id string, input model.KeretaInput) model.Kereta {
	return n.Method.UpdateTrain(id, input)
}

func (n *KeretaService) Delete(id string) bool {
	return n.Method.DeleteTrain(id)
}
