package controllers

import (
	"github.com/chaitanyamaili/financial-account-management/api/models"
	"github.com/chaitanyamaili/financial-account-management/api/services"
	"github.com/gofiber/fiber/v2"
)

func CreateAccount(c *fiber.Ctx) error {
	account := new(models.Account)
	if err := c.BodyParser(account); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if err := services.CreateAccount(account); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create account"})
	}

	return c.Status(fiber.StatusCreated).JSON(account)
}

func GetAllAccounts(c *fiber.Ctx) error {
	accounts, err := services.GetAllAccounts()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to fetch accounts"})
	}
	return c.JSON(accounts)
}

// TODO
// GetAccountByUserID
// GetAccountByName
