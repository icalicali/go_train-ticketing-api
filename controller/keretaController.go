package controller

import (
	"go_mini-project/model"
	"go_mini-project/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

var keretaService service.KeretaService = service.NewTrain()

func GetAllTrain(c echo.Context) error {
	var kereta []model.Kereta = keretaService.GetAll()

	return c.JSON(http.StatusOK, kereta)
}

func GetTrainByID(c echo.Context) error {
	var id string = c.Param("id")

	kereta := keretaService.GetByID(id)

	if kereta.ID == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "ticket not found",
		})
	}

	return c.JSON(http.StatusOK, kereta)
}

func CreateTrain(c echo.Context) error {
	var input *model.KeretaInput = new(model.KeretaInput)

	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request",
		})
	}

	err := input.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "validasi gagal",
		})
	}

	kereta := keretaService.Create(*input)

	return c.JSON(http.StatusCreated, kereta)
}

func UpdateTrain(c echo.Context) error {
	var input *model.KeretaInput = new(model.KeretaInput)

	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request",
		})
	}

	var keretaId string = c.Param("id")

	err := input.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "validasi gagal",
		})
	}

	kereta := keretaService.Update(keretaId, *input)

	return c.JSON(http.StatusOK, kereta)
}

func DeleteTrain(c echo.Context) error {
	var keretaId string = c.Param("id")

	isSuccess := keretaService.Delete(keretaId)

	if !isSuccess {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "failed to delete data",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "data deleted",
	})

}
