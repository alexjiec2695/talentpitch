package restchallenges

import "github.com/gin-gonic/gin"

func Handler(client *gin.Engine, controller Controller) {
	client.GET("/challenges", controller.getChallenges)
	client.POST("/challenges", controller.createChallenges)
	client.GET("/challenges/:id", controller.getChallengesByID)
	client.PUT("/challenges/:id", controller.updateChallenges)
	client.DELETE("/challenges/:id", controller.deleteChallenges)
}
