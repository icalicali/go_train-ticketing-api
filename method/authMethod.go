package method

import (
	"go_mini-project/database"
	"go_mini-project/jwt"
	"go_mini-project/model"

	"golang.org/x/crypto/bcrypt"
)

type AuthenticationMethod interface {
	Register(input model.Register) model.Pemesan
	Login(input model.Register) string
}

type AuthenticationSystem struct{}

func (a *AuthenticationSystem) Register(input model.Register) model.Pemesan {
	password, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	var newUser model.Pemesan = model.Pemesan{
		Username: input.Username,
		Password: string(password),
	}

	var createdUser model.Pemesan = model.Pemesan{}

	result := database.DB.Create(&newUser)

	result.Last(&createdUser)

	return createdUser
}

func (a *AuthenticationSystem) Login(input model.Register) string {
	var user model.Pemesan = model.Pemesan{}

	database.DB.First(&user, "username = ?", input.Username)

	if user.ID == 0 {
		return ""
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))

	if err != nil {
		return ""
	}

	token := jwt.CreateToken(user.ID)

	return token
}
