package controllers

import (
	"GoBigQuery/config"
	"GoBigQuery/models"
	"fmt"

	"github.com/google/uuid"
)

func CreateHotel(hotel *models.Hotel) {
	config.DB.Create(&hotel)
}

func GetHotelByID(HotelID uuid.UUID) (*models.Hotel, error) {
	var Hotel models.Hotel
	result := config.DB.First(&Hotel, "id = ?", HotelID)
	return &Hotel, result.Error
}

func GetAllHotels() ([]models.Hotel, error) {
	var hotels []models.Hotel
	result := config.DB.Find(&hotels)
	return hotels, result.Error
}

func UpdateHotel(hotel *models.Hotel, hotelID uuid.UUID) {
	fmt.Printf("%+v\n", hotel.Model)
	config.DB.Model(&models.Hotel{}).Where("id = ?", hotelID).Updates(*hotel)
}

func DeleteHotel(hotelID uuid.UUID) {
	config.DB.Delete(&models.Hotel{}, "id = ?", hotelID)
}
