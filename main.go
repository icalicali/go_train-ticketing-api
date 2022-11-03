package main

import (
	"go_mini-project/database"
	"go_mini-project/route"

	"github.com/labstack/echo/v4"
)

func main() {
	database.Connect()

	server := echo.New()

	route.SetupRoute(server)

	server.Logger.Fatal(server.Start(":1323"))
}
