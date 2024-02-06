package routes

import (
	"net/http"
	"strconv"
	"udemy/src/repository"

	"github.com/labstack/echo/v4"
)

// InitRoutes initializes all the routes for the application
func InitRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		advertisement := repository.AdvertisementRepository{}
		return c.JSONPretty(http.StatusCreated, advertisement.List(), "  ")
	})

	e.GET("/:longitude/:latitude", func(c echo.Context) error {
		longitudeStr := c.Param("longitude")
		latitudeStr := c.Param("latitude")

		longitude, err := strconv.ParseFloat(longitudeStr, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid longitude"})
		}
		latitude, err := strconv.ParseFloat(latitudeStr, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid latitude"})
		}

		advertisement := repository.AdvertisementRepository{}
		return c.JSONPretty(http.StatusCreated, advertisement.ListByGeoLocation(longitude, latitude), "  ")
	})

	// Add more routes here
}
