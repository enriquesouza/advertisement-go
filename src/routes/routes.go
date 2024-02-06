package routes

import (
	"net/http"

	_ "advertisement-go/docs"
	"advertisement-go/src/repository"
	"advertisement-go/src/routes/handlers"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func InitRoutes(e *echo.Echo) {

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.GET("/", func(c echo.Context) error {
		advertisement := repository.AdvertisementRepository{}
		return c.JSONPretty(http.StatusCreated, advertisement.List(), "  ")
	})

	e.GET("/:longitude/:latitude", handlers.GeoSearchHandler)

	e.POST("/search", handlers.GeoSearchWithMaxDistanceHandler)
}
