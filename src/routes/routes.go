package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"
	_ "udemy/docs"
	"udemy/src/repository"
	"udemy/src/routes/handlers"
)

func InitRoutes(e *echo.Echo) {

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("/", func(c echo.Context) error {
		advertisement := repository.AdvertisementRepository{}
		return c.JSONPretty(http.StatusCreated, advertisement.List(), "  ")
	})
	// GeoSearchHandler searches for advertisements based on geographic location.
	// @Summary Search by geo location
	// @Description Searches for advertisements by longitude and latitude.
	// @Tags advertisement
	// @Accept  json
	// @Produce  json
	// @Param   longitude    path      float64  true  "Longitude"
	// @Param   latitude     path      float64  true  "Latitude"
	// @Success 200 {array}  Advertisement  "A list of advertisements"
	// @Router /{longitude}/{latitude} [get]
	//http://localhost:1323/-46.6773326/-23.5800755
	e.GET("/:longitude/:latitude", handlers.GeoSearchHandler)

	// GeoSearchWithMaxDistanceHandler searches for advertisements based on geographic location and maximum distance.
	// @Summary Search by geo location with max distance
	// @Description Searches for advertisements by longitude, latitude, and maximum distance.
	// @Tags advertisement
	// @Accept  json
	// @Produce  json
	// @Param   geoLocationRequest  body      GeoLocationRequest  true  "Geo Location Request"
	// @Success 200 {array}  Advertisement  "A list of advertisements within max distance"
	// @Router /search [post]
	e.POST("/search", handlers.GeoSearchWithMaxDistanceHandler)
}
