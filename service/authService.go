package service

import (
	"go_mini-project/method"
	"go_mini-project/model"
)

//authSvc

type AuthService struct {
	Method method.AuthenticationMethod
}

func NewAuthService() AuthService {
	return AuthService{
		Method: &method.AuthenticationSystem{},
	}
}

func (a *AuthService) Register(input model.Register) model.Pemesan {
	return a.Method.Register(input)
}

func (a *AuthService) Login(input model.Register) string {
	return a.Method.Login(input)
}
