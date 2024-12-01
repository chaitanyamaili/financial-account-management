package main

import (
	"log"

	"github.com/chaitanyamaili/financial-account-management/api/config"
	"github.com/chaitanyamaili/financial-account-management/api/database"
	"github.com/chaitanyamaili/financial-account-management/api/models"
	"github.com/chaitanyamaili/financial-account-management/api/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.LoadConfig()

	// Connect to the database
	database.ConnectDB(cfg)

	// Migrate the database
	database.DB.AutoMigrate(&models.User{}, &models.Account{}, &models.Transaction{})

	// Create a new Fiber instance
	app := fiber.New()

	// Setup routes
	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":8080"))
}
