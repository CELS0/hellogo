package routes

import (
	"github.com/CELS0/hellogo/controllers"
	"github.com/CELS0/hellogo/middlewares"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		login := main.Group("login")
		{
			login.POST("/", controllers.Login)
		}
		user := main.Group("user")
		{
			user.POST("/", controllers.CreateUser)
			user.GET("/:id", controllers.ShowUser)
			user.GET("/", controllers.ShowUsers)
			user.PUT("/", controllers.UpdateUser)
			user.DELETE("/:id", controllers.DeleteUser)
		}

		books := main.Group("books", middlewares.Auth())
		{
			books.GET("/:id", controllers.ShowBook)
			books.GET("/", controllers.ShowBooks)
			books.POST("/", controllers.CreateBook)
			books.PUT("/", controllers.UpdateBook)
			books.DELETE("/:id", controllers.DeleteBook)
		}
	}
	return router
}
