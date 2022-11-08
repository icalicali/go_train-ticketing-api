package method

import "go_mini-project/model"

type TiketMethod interface {
	GetAll() []model.Tiket
	GetByID(id string) model.Tiket
	Create(input model.TiketInput) model.Tiket
	Update(id string, input model.TiketInput) model.Tiket
	Delete(id string) bool
}

type AuthenticationMethod interface {
	Register(input model.Register) model.Pemesan
	Login(input model.Register) string
}
