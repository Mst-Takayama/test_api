package routes

import (
	"test_api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {
	userRoutes := router.Group("/users")
	{
		userRoutes.GET("/", controllers.GetUsers)
	}
}
