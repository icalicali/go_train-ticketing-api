package route

import (
	"go_mini-project/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupRoute(server *echo.Echo) {
	server.POST("/register", controller.Register)
	server.POST("/login", controller.Login)

	ticketRoute := server.Group("")

	ticketRoute.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("SECRET_KEY"),
	}))

	ticketRoute.GET("/tiket", controller.GetAll)
	ticketRoute.GET("/tiket/:id", controller.GetByID)
	ticketRoute.POST("/tiket", controller.Create)
	ticketRoute.PUT("/tiket/:id", controller.Update)
	ticketRoute.DELETE("/tiket/:id", controller.Delete)

	trainRoute := server.Group("")

	trainRoute.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("SECRET_KEY"),
	}))

	trainRoute.GET("/kereta", controller.GetAllTrain)
	trainRoute.GET("/kereta/:id", controller.GetTrainByID)
	trainRoute.POST("/kereta", controller.CreateTrain)
	trainRoute.PUT("/kereta/:id", controller.UpdateTrain)
	trainRoute.DELETE("/kereta/:id", controller.DeleteTrain)
}
