package services

import (
	"github.com/chaitanyamaili/financial-account-management/api/database"
	"github.com/chaitanyamaili/financial-account-management/api/models"
)

func CreateAccount(account *models.Account) error {
	return database.DB.Create(account).Error
}

func GetAllAccounts() ([]models.Account, error) {
	var accounts []models.Account
	err := database.DB.Find(&accounts).Error
	return accounts, err
}

func GetAccountByUserID(userID uint) ([]models.Account, error) {
	var accounts []models.Account
	err := database.DB.Where("user_id = ?", userID).Find(&accounts).Error
	return accounts, err
}

func GetAccountByName(name string) (models.Account, error) {
	var account models.Account
	err := database.DB.Where("account_name = ?", name).First(&account).Error
	return account, err
}
