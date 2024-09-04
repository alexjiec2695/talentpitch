package restvideo

import "github.com/gin-gonic/gin"

func Handler(client *gin.Engine, controller Controller) {
	client.GET("/videos", controller.getVideo)
	client.POST("/videos", controller.createVideo)
	client.GET("/videos/:id", controller.getVideoByID)
	client.PUT("/videos/:id", controller.updateVideo)
	client.DELETE("/videos/:id", controller.deleteVideo)
}
