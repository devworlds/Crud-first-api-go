package routes

import (
	controllers "api/controllers"

	"github.com/gin-gonic/gin"
)
func AppRoutes(router *gin.Engine) *gin.RouterGroup {
	tweetController := controllers.NewTweetController()

	v1 := router.Group("/v1")
	{
		v1.GET("/tweet", tweetController.FindAll)
		v1.POST("/tweet", tweetController.Create)
		v1.DELETE("/tweet/:id", tweetController.Delete)
		v1.GET("/tweet/:id", tweetController.FindOne)
		v1.PATCH("/tweet/:id", tweetController.Update)
	}

	return v1
}