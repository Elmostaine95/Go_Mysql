package routes

import (
	"GoBigQuery/controllers"
	"GoBigQuery/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func SetupRoomRoutes(router *gin.Engine) {
	roomGroup := router.Group("/rooms")
	{
		roomGroup.POST("/", createRoom)
		roomGroup.GET("/:id", getRoomByID)
		roomGroup.GET("/", getAllRooms)
		roomGroup.PUT("/:id", updateRoom)
		roomGroup.DELETE("/:id", deleteRoom)
	}
}

func createRoom(c *gin.Context) {
	var room models.Room
	c.BindJSON(&room)
	room.ID = uuid.New()
	controllers.CreateRoom(&room)
	c.JSON(201, room)
}

func getRoomByID(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	room, err := controllers.GetRoomByID(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "Room not found"})
		return
	}
	c.JSON(200, room)
}

func getAllRooms(c *gin.Context) {
	rooms, _ := controllers.GetAllRooms()
	c.JSON(200, rooms)
}

func updateRoom(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	var room models.Room
	c.BindJSON(&room)
	controllers.UpdateRoom(&room,id)
	c.JSON(200, room)
}

func deleteRoom(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	controllers.DeleteRoom(id)
	c.JSON(204, nil)
}
