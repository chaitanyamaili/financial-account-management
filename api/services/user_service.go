package services

import (
	"github.com/chaitanyamaili/financial-account-management/api/database"
	"github.com/chaitanyamaili/financial-account-management/api/models"
)

// CreateUser creates a new user in the database
// using the provided User object
// and returns an error if any
func CreateUser(user *models.User) error {
	return database.DB.Create(user).Error
}

// GetAllUsers fetches all users from the database
// and returns them as a slice of User objects
// along with an error if any
func GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := database.DB.Find(&users).Error
	return users, err
}

func DeleteUserByUsername(username string) error {
	return database.DB.Where("username = ?", username).Delete(&models.User{}).Error
}
