package main

import (
	"GoBigQuery/config"
	"GoBigQuery/models"
	"GoBigQuery/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	db := config.DB

	// AutoMigrate will create the "users" table based on the User struct
	db.AutoMigrate(&models.Hotel{})
	db.AutoMigrate(&models.Reservation{})
	db.AutoMigrate(&models.Room{})
	db.AutoMigrate(&models.User{})

	r := gin.Default()

	routes.SetupHotelRoutes(r)
	routes.SetupReservationRoutes(r)
	routes.SetupRoomRoutes(r)
	routes.SetupUserRoutes(r)

	r.Run(":8080")
}
