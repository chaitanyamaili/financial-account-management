package controllers

import (
	"github.com/chaitanyamaili/financial-account-management/api/models"
	"github.com/chaitanyamaili/financial-account-management/api/services"
	"github.com/gofiber/fiber/v2"
)

func CreateTransaction(c *fiber.Ctx) error {
	transaction := new(models.Transaction)
	if err := c.BodyParser(transaction); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := services.CreateTransaction(transaction); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create transaction"})
	}

	return c.Status(fiber.StatusCreated).JSON(transaction)
}
