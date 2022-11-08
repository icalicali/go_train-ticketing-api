package controller

import (
	"go_mini-project/model"
	"go_mini-project/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

var tiketService service.TiketService = service.New()

func GetAll(c echo.Context) error {
	var tikets []model.Tiket = tiketService.GetAll()

	// melakukan parsing data dalam bentuk JSON
	return c.JSON(http.StatusOK, tikets)
}

func GetByID(c echo.Context) error {
	var id string = c.Param("id")

	tiket := tiketService.GetByID(id)

	// jika data tidak ditemukan
	if tiket.ID == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "ticket not found",
		})
	}

	return c.JSON(http.StatusOK, tiket)
}

func Create(c echo.Context) error {
	var input *model.TiketInput = new(model.TiketInput)

	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request",
		})
	}

	tiket := tiketService.Create(*input)

	return c.JSON(http.StatusCreated, tiket)
}

func Update(c echo.Context) error {
	var input *model.TiketInput = new(model.TiketInput)

	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request",
		})
	}

	var tiketId string = c.Param("id")

	tiket := tiketService.Update(tiketId, *input)

	return c.JSON(http.StatusOK, tiket)
}

func Delete(c echo.Context) error {
	var tiketId string = c.Param("id")

	isSuccess := tiketService.Delete(tiketId)

	if !isSuccess {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "failed to delete data",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "data deleted",
	})

}
