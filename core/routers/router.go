package routers

import (
	"github.com/akposiyefa/go-gin-migration/core/handlers"
	"github.com/gin-gonic/gin"
)

func ApiRoutes(router *gin.Engine) {
	subRoute := router.Group("api/v1")
	subRoute.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to my simple api with gin",
			"data": map[string]string{
				"Name":  "Orutu Akposieyefa Williams",
				"Email": "orutu1@gmail.com",
				"Role":  "Software Engineer",
			},
			"success": true,
		})
	})

	//user routers
	usersRoute := subRoute.Group("/users")
	usersRoute.POST("/create", handlers.CreateUser)
	usersRoute.GET("/all", handlers.GetUsers)
	usersRoute.GET("/single/:id", handlers.GetUser)
	usersRoute.PATCH("/update/:id", handlers.UpdateUser)
	usersRoute.DELETE("/delete/:id", handlers.DeleteUser)
}
