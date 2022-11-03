package route

import (
	"go_mini-project/controller"

	"github.com/labstack/echo/v4"
)

func SetupRoute(server *echo.Echo) {
	server.GET("/mini/kai/tiket", controller.GetAll)
	server.GET("/mini/kai/tiket/:id", controller.GetByID)
	server.POST("/mini/kai/tiket", controller.Create)
	server.PUT("/mini/kai/tiket/:id", controller.Update)
	server.DELETE("/mini/kai/tiket/:id", controller.Delete)
}
