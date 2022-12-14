package main

import (
	"TerraInnAPI/database"
	"TerraInnAPI/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"flag"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PATCH,PUT,DELETE",
		AllowCredentials: true,
	}))

	// Static folder
	app.Static("/assets", "./assets")

	// Connect Database
	if !database.Connect() {
		return
	}

	// Init router
	routes.InitRoutes(app)

	// Run app
	// port := config.Config("ENV_PORT")
	port := "3000"

	addr := flag.String("addr", ":"+port, "http service address")
	flag.Parse()
	log.Fatal(app.Listen(*addr))
}
