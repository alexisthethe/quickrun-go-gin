package controllers

import (
	"github.com/alexisthethe/quickrun-go-gin/initializers"
	"github.com/alexisthethe/quickrun-go-gin/models"
	"github.com/gin-gonic/gin"
)

func CreateArticle (c *gin.Context) {

	var requestBody struct {
		Content string
		Title string
	}

	c.Bind(&requestBody)

	article := models.Article{Title: requestBody.Title, Content: requestBody.Content}

	result := initializers.DB.Create(&article)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, article)
}

func GetAllArticle (c *gin.Context) {
	var articles []models.Article
	initializers.DB.Find(&articles)

	c.JSON(200, articles)
}

func GetOneArticle (c *gin.Context) {
	id := c.Param("id")

	var article models.Article
	result := initializers.DB.First(&article, id)
	if result.Error != nil {
		c.Status(404)
		return
	}

	c.JSON(200, article)
}

func UpdateArticle (c *gin.Context) {
	id := c.Param("id")

	var requestBody struct {
		Content string
		Title string
	}

	c.Bind(&requestBody)

	var article models.Article
	result := initializers.DB.First(&article, id)
	if result.Error != nil {
		c.Status(404)
		return
	}

	initializers.DB.Model(&article).Updates(models.Article{Title: requestBody.Title, Content: requestBody.Content})

	c.JSON(200, article)
}

func DeleteArticle(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Article{}, id)

	c.Status(200)
}