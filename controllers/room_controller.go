package controllers

import (
	"GoBigQuery/config"
	"GoBigQuery/models"

	"github.com/google/uuid"
)

func CreateRoom(room *models.Room) {
	config.DB.Create(&room)
}

func GetRoomByID(RoomID uuid.UUID) (*models.Room, error) {
	var room models.Room
	result := config.DB.First(&room, "id = ?", RoomID)
	return &room, result.Error
}

func GetAllRooms() ([]models.Room, error) {
	var rooms []models.Room
	result := config.DB.Joins("Hotel").Find(&rooms)
	return rooms, result.Error
}

func UpdateRoom(Room *models.Room, RoomID uuid.UUID) {
	config.DB.Model(&models.Room{}).Where("id = ?", RoomID).Updates(*Room)
}

func DeleteRoom(RoomID uuid.UUID) {
	config.DB.Delete(&models.Room{}, "id = ?", RoomID)
}
