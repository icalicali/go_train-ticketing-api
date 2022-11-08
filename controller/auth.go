package controller

import (
	"go_mini-project/model"
	"go_mini-project/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

var authService service.AuthService = service.NewAuthService()

func Register(c echo.Context) error {
	var Register *model.Register = new(model.Register)

	if err := c.Bind(Register); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "REQUEST ERROR",
		})
	}

	err := Register.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "VALIDATION ERROR",
		})
	}

	user := authService.Register(*Register)

	return c.JSON(http.StatusCreated, user)

}

func Login(c echo.Context) error {
	var Register *model.Register = new(model.Register)

	if err := c.Bind(Register); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "REQUEST ERROR",
		})
	}

	err := Register.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "VALIDATION ERROR",
		})
	}

	token := authService.Login(*Register)

	if token == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "LOGIN ERROR. Please try again soon.",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})

}
