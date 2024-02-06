package main

import (
	"log"
	"udemy/src/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func init() {
	// Load the .env file
	err := godotenv.Load() // Tries to load from the current directory
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func main() {
	e := echo.New()
	
	routes.InitRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
