package routes

import (
	"github.com/CELS0/hellogo/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		books := main.Group("books")
		{
			books.GET("/", controllers.ShowBook)
		}
	}
	return router
}
