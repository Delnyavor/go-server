package routes

import (
	"github.com/Delnyavor/go-fiber-mongo-hrms/controllers"
	"github.com/gofiber/fiber/v2"
)

func RegisterEmployeeRoutes(app *fiber.App) {

	app.Get("/employees", controllers.GetEmployees)
	app.Post("/employee", controllers.CreateEmployee)
	app.Put("/employee/:id", controllers.UpdateEmployee)
	app.Delete("/employee/:id", controllers.DeleteEmployee)
}
