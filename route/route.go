package route

import (
	"go_mini-project/controller"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupRoute(server *echo.Echo) {
	server.POST("/mini/kai/pemesan/register", controller.Register)
	server.POST("/mini/kai/pemesan/login", controller.Login)

	privateRoutes := server.Group("")

	privateRoutes.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("SECRET_KEY"),
	}))

	// TEST
	privateRoutes.GET("/test", func(c echo.Context) error {
		return c.String(http.StatusOK, "Authentication Success!")
	})

	privateRoutes.GET("/mini/kai/tiket", controller.GetAll)
	privateRoutes.GET("/mini/kai/tiket/:id", controller.GetByID)
	privateRoutes.POST("/mini/kai/tiket", controller.Create)
	privateRoutes.PUT("/mini/kai/tiket/:id", controller.Update)
	privateRoutes.DELETE("/mini/kai/tiket/:id", controller.Delete)
}
