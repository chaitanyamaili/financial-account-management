package routes

import (
	"github.com/chaitanyamaili/financial-account-management/api/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// User routes
	api.Post("/users", controllers.CreateUser)
	api.Get("/users", controllers.GetAllUsers)
}
