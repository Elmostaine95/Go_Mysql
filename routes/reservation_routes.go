package routes

import (
	"GoBigQuery/controllers"
	"GoBigQuery/models"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func SetupReservationRoutes(router *gin.Engine) {
	reservationGroup := router.Group("/reservations")
	{
		reservationGroup.POST("/", createReservation)
		reservationGroup.GET("/:id", getReservationByID)
		reservationGroup.GET("/", getAllReservations)
		reservationGroup.GET("/getReservetionByUserid/:id", getReservetionByUserID)
		reservationGroup.GET("/getHotelRevenueByDate/:hotelID", getHotelRevenueByDate)
		reservationGroup.PUT("/:id", updateReservation)
		reservationGroup.DELETE("/:id", deleteReservation)
	}
}

func createReservation(c *gin.Context) {
	var reservation models.Reservation
	c.BindJSON(&reservation)
	reservation.ID = uuid.New()
	controllers.CreateReservation(&reservation)
	c.JSON(201, reservation)
}

func getReservationByID(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	reservation, err := controllers.GetReservationByID(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "Reservation not found"})
		return
	}
	c.JSON(200, reservation)
}

func getAllReservations(c *gin.Context) {
	reservations, _ := controllers.GetAllReservations()
	c.JSON(200, reservations)
}

func updateReservation(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	var reservation models.Reservation
	c.BindJSON(&reservation)
	controllers.UpdateReservation(&reservation, id)
	c.JSON(200, reservation)
}

func deleteReservation(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	controllers.DeleteReservation(id)
	c.JSON(204, nil)
}

func getReservetionByUserID(c *gin.Context) {
	user_id, _ := uuid.Parse(c.Param("id"))
	reservation, err := controllers.GetReservetionByUserID(user_id)
	if err != nil {
		c.JSON(404, gin.H{"error": "Reservation not found"})
		return
	}
	c.JSON(200, reservation)
}

func getHotelRevenueByDate(c *gin.Context) {
	hotel_id, _ := uuid.Parse(c.Param("id"))
	const layout = "2006-01-02T15:04:05Z"
	fromDate, _ := time.Parse(layout, c.Query("FromDate"))
	toDate, _ := time.Parse(layout, c.Query("ToDate"))

	//Date validation
	if fromDate.After(toDate) {
		c.JSON(404, gin.H{"error": "Invalid Dates"})
		return
	}

	fmt.Printf("from date %+v to date %+v\n", fromDate, toDate)
	reservation := controllers.GetHotelRevenueByDate(hotel_id, fromDate, toDate)
	c.JSON(200, reservation)
}
