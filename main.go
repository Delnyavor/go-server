package main

import (
	"log"

	"github.com/Delnyavor/go-fiber-mongo-hrms/database"
	"github.com/Delnyavor/go-fiber-mongo-hrms/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}
	app := fiber.New()
	routes.RegisterEmployeeRoutes(app)
	app.Listen(":3000")

	// defer database.Disconnect()

}
