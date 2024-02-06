package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	_ "udemy/src/domain"
	"udemy/src/repository"
)

type GeoLocationRequest struct {
	Longitude   float64 `json:"longitude"`
	Latitude    float64 `json:"latitude"`
	MaxDistance int32   `json:"maxDistance"` // In Kilometers, fixed typo in struct tag
}

// GeoSearchHandler searches for advertisements based on geographic location.
// @Summary Search by geo location
// @Description Searches for advertisements by longitude and latitude.
// @Tags advertisement
// @Accept  json
// @Produce  json
// @Param   longitude    path      float64  true  "Longitude"
// @Param   latitude     path      float64  true  "Latitude"
// @Success 200 {array}  domain.Advertisement  "A list of advertisements"
// @Router /{longitude}/{latitude} [get]
func GeoSearchHandler(c echo.Context) error {
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
}

// GeoSearchWithMaxDistanceHandler searches for advertisements based on geographic location and maximum distance.
// @Summary Search by geo location with max distance
// @Description Searches for advertisements by longitude, latitude, and maximum distance.
// @Tags advertisement
// @Accept  json
// @Produce  json
// @Param   geoLocationRequest  body      GeoLocationRequest  true  "Geo Location Request"
// @Success 200 {array}  domain.Advertisement  "A list of advertisements within max distance"
// @Router /search [post]
func GeoSearchWithMaxDistanceHandler(c echo.Context) error {
	var req GeoLocationRequest

	// Bind the JSON body into the req struct
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request body"})
	}

	// Now, req.Longitude and req.Latitude have the values
	var advertisement = repository.AdvertisementRepository{}
	return c.JSONPretty(http.StatusCreated, advertisement.ListByGeoLocationWithMaxDistance(req.Longitude, req.Latitude, req.MaxDistance), "  ")
}
