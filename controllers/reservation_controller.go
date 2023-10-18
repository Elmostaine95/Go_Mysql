package controllers

import (
	"GoBigQuery/config"
	"GoBigQuery/models"
	"time"

	"github.com/google/uuid"
)

func CreateReservation(reservation *models.Reservation) {
	config.DB.Create(&reservation)
}

func GetReservationByID(ReservationID uuid.UUID) (*models.Reservation, error) {
	var reservation models.Reservation
	result := config.DB.First(&reservation, "id = ?", ReservationID)
	return &reservation, result.Error
}

func GetAllReservations() ([]models.Reservation, error) {
	var reservations []models.Reservation
	result := config.DB.Preload("Hotel").Preload("Room.Hotel").Preload("User").Find(&reservations)
	return reservations, result.Error
}

func UpdateReservation(reservation *models.Reservation, reservationID uuid.UUID) {
	config.DB.Model(&models.Reservation{}).Where("id = ?", reservationID).Updates(*reservation)
}

func DeleteReservation(reservationsID uuid.UUID) {
	config.DB.Delete(&models.Reservation{}, "id = ?", reservationsID)
}

func GetReservetionByUserID(UserID uuid.UUID) ([]models.Reservation, error) {
	var reservations []models.Reservation
	result := config.DB.Preload("Hotel").Preload("Room.Hotel").Where("user_id = ?", UserID).Find(&reservations)
	return reservations, result.Error
}

func GetHotelRevenueByDate(UserID uuid.UUID, fromDate time.Time, toDate time.Time) int {
	var sum int
	config.DB.Table("reservations").
		Preload("Room.Hotel").
		Select("sum(CASE WHEN check_out > ? THEN DATEDIFF(?, check_in) ELSE DATEDIFF(check_out, check_in) )", toDate, toDate).
		Where("check_in >= ?", fromDate).
		Scan(&sum)

	return sum
}
