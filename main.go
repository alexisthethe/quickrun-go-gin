package main

import (
	"github.com/alexisthethe/quickrun-go-gin/controllers"
	"github.com/alexisthethe/quickrun-go-gin/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	// Load .env file
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main() {
	r := gin.Default()
	r.SetTrustedProxies(nil)

	r.POST("/articles", controllers.CreateArticle)
	r.GET("/articles", controllers.GetAllArticle)
	r.GET("/articles/:id", controllers.GetOneArticle)
	r.PUT("/articles/:id", controllers.UpdateArticle)
	r.DELETE("/articles/:id", controllers.DeleteArticle)
	
	r.Run()
}