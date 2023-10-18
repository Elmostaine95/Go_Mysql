package controllers

import (
	"GoBigQuery/config"
	"GoBigQuery/models"

	"github.com/google/uuid"
)

func CreateUser(user *models.User) {
	config.DB.Create(&user)
}

func GetUserByID(userID uuid.UUID) (*models.User, error) {
	var user models.User
	result := config.DB.First(&user, "id = ?", userID)
	return &user, result.Error
}

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := config.DB.Find(&users)
	return users, result.Error
}

func UpdateUser(User *models.User, UserID uuid.UUID) {
	config.DB.Model(&models.User{}).Where("id = ?", UserID).Updates(*User)
}

func DeleteUser(userID uuid.UUID) {
	config.DB.Delete(&models.User{}, "id = ?", userID)
}
