package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
    e.POST("/create", create)
    e.POST("/get", get)
	e.Logger.Fatal(e.Start(":1323"))
}

func create(c echo.Context) error {
	return c.String(http.StatusOK, "Workout created")
}

func get(c echo.Context) error {
	return c.JSON(http.StatusOK, "Getting workout")
}
