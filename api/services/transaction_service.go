package services

import (
	"github.com/chaitanyamaili/financial-account-management/api/database"
	"github.com/chaitanyamaili/financial-account-management/api/models"
)

func CreateTransaction(transaction *models.Transaction) error {
	return database.DB.Create(transaction).Error
}

func GetAllTransactions() ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := database.DB.Find(&transactions).Error
	return transactions, err
}

func GetTransactionByAccountID(accountID uint) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := database.DB.Where("account_id = ?", accountID).Find(&transactions).Error
	return transactions, err
}
