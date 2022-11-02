package method

import "go_mini-project/model"

type TiketMethod interface {
	GetAll() []model.Tiket
	GetByID(id string) model.Tiket
	Create(input model.TiketInput) model.Tiket
	Update(id string, input model.TiketInput) model.Tiket
	Delete(id string) bool
}
