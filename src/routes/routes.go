package routes

import (
	"net/http"
	"udemy/src/repository"

	"github.com/labstack/echo/v4"
)

// InitRoutes initializes all the routes for the application
func InitRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		advertisement := repository.AdvertisementRepository{}
		return c.JSONPretty(http.StatusCreated, advertisement.List(), "  ")
	})

	// Add more routes here
}
