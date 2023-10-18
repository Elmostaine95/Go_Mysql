package routes

import (
	"GoBigQuery/controllers"
	"GoBigQuery/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func SetupUserRoutes(router *gin.Engine) {
	userGroup := router.Group("/users")
	{
		userGroup.POST("/", createUser)
		userGroup.GET("/:id", getUserByID)
		userGroup.GET("/", getAllUsers)
		userGroup.PUT("/:id", updateUser)
		userGroup.DELETE("/:id", deleteUser)
	}
}

func createUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	user.ID = uuid.New()
	controllers.CreateUser(&user)
	c.JSON(201, user)
}

func getUserByID(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	user, err := controllers.GetUserByID(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}
	c.JSON(200, user)
}

func getAllUsers(c *gin.Context) {
	users, _ := controllers.GetAllUsers()
	c.JSON(200, users)
}

func updateUser(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	var user models.User
	c.BindJSON(&user)
	controllers.UpdateUser(&user, id)
	c.JSON(200, user)
}

func deleteUser(c *gin.Context) {
	id, _ := uuid.Parse(c.Param("id"))
	controllers.DeleteUser(id)
	c.JSON(204, nil)
}
