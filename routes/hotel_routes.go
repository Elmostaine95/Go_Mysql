package routes

import (
	"GoBigQuery/controllers"
	"GoBigQuery/models"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

func SetupHotelRoutes(router *gin.Engine) {
	hotelGroup := router.Group("/hotels")
	{
		hotelGroup.POST("/", createHotel)
		hotelGroup.GET("/:id", getHotelByID)
		hotelGroup.GET("/", getAllHotels)
		hotelGroup.PUT("/:id", updateHotel)
		hotelGroup.DELETE("/:id", deleteHotel)
	}
}

func createHotel(c *gin.Context) {
	var hotel models.Hotel
	c.BindJSON(&hotel)
	hotel.ID = uuid.New()
	controllers.CreateHotel(&hotel)
	c.JSON(201, hotel)
}

func getHotelByID(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	hotel, err := controllers.GetHotelByID(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "Hotel not found"})
		return
	}
	c.JSON(200, hotel)
}

func getAllHotels(c *gin.Context) {
	hotels, _ := controllers.GetAllHotels()
	c.JSON(200, hotels)
}

func updateHotel(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	var hotel models.Hotel
	c.BindJSON(&hotel)
	controllers.UpdateHotel(&hotel, id)
	c.JSON(200, hotel)
}

func deleteHotel(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	controllers.DeleteHotel(id)
	c.JSON(204, nil)
}
