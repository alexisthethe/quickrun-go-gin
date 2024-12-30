package main

import (
	"github.com/alexisthethe/quickrun-go-gin/initializers"
	"github.com/alexisthethe/quickrun-go-gin/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}

func main() {
	initializers.DB.AutoMigrate(&models.Article{})
}