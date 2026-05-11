
package main

import (
	"log"
	"net/http"

	"github.com/77jennifer77/workout-management-service/internal/workout"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	e := echo.New()
	e.POST("/create", create)
	e.POST("/get", create)
	e.Logger.Fatal(e.Start(":1323"))
}

func create(c echo.Context) error {
	workout := workout.Workout{}
	if err := c.Bind(&workout); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err := workout.CreateWorkout()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, "Workout created")
}

func get(c echo.Context) error {
	workout := workout.Workout{}
	if err := c.Bind(&workout); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	err := workout.GetWorkout()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, workout)
}
