package main

import (
	"go-restapi/controllers"
	"go-restapi/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.DBConnection()
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello from FINAL DEMO presentation (June 8)!!!",
		})
	})

	r.GET("/posts", controllers.GetAllPosts)
	r.GET("/posts/:id", controllers.GetPost)
	r.POST("/posts", controllers.CreatePost)
	r.PUT("/posts/:id", controllers.UpdatePost)
	r.DELETE("/posts/:id", controllers.DeletePost)

	return r
}

func main() {
	r := setupRouter()
	r.Run(":9090")
}
