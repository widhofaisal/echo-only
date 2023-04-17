package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World! Perjalanan golang saya dimulai dari sini")
	})

	e.GET("/users", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "mengakses route /users",
		})
	})

	e.Logger.Fatal(e.Start(":1323"))
}
