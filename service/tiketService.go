package service

import (
	"go_mini-project/method"
	"go_mini-project/model"
)

//ticketSvc

type TiketService struct {
	Method method.TiketMethod
}

func New() TiketService {
	return TiketService{
		Method: &method.TiketMethodCRUD{},
	}
}

func (n *TiketService) GetAll() []model.Tiket {
	return n.Method.GetAll()
}

func (n *TiketService) GetByID(id string) model.Tiket {
	return n.Method.GetByID(id)
}

func (n *TiketService) Create(input model.TiketInput) model.Tiket {
	return n.Method.Create(input)
}

func (n *TiketService) Update(id string, input model.TiketInput) model.Tiket {
	return n.Method.Update(id, input)
}

func (n *TiketService) Delete(id string) bool {
	return n.Method.Delete(id)
}
