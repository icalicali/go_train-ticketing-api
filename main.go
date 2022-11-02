package main

import (
	"go_mini-project/database"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	database.Connect()

	// membuat instance echo
	e := echo.New()

	// mendaftar route "/" dengan method get
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// menjalankan server
	e.Logger.Fatal(e.Start(":1323"))
}
