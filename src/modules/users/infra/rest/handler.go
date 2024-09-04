package restuser

import "github.com/gin-gonic/gin"

func Handler(client *gin.Engine, controller Controller) {
	client.GET("/users", controller.getUsers)
	client.POST("/users", controller.createUser)
	client.GET("/users/:id", controller.getUserByID)
	client.PUT("/users/:id", controller.updateUser)
	client.DELETE("/users/:id", controller.deleteUser)
}
