package controllers

import (
	"fmt"
	"go-restapi/initializers"
	"go-restapi/models"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	// Get data off req Body
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	// Create a post

	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	// Throw error
	if result.Error != nil {
		c.Status(400)
		return
	}
	// Return it
	c.JSON(201, gin.H{
		"post": post,
	})
}

func GetAllPosts(c *gin.Context) {
	// Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// Respond with them
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func GetPost(c *gin.Context) {
	// Get id off url
	id := c.Param("id")
	// Get the post
	var post models.Post
	result := initializers.DB.First(&post, id)

	// Throw error
	if result.Error != nil {
		response := fmt.Sprintf("the post with id %s does not exists.", id)
		c.JSON(404, gin.H{
			"message": response,
		})
		return
	}
	// Respond with them
	c.JSON(200, gin.H{
		"posts": post,
	})
}

func UpdatePost(c *gin.Context) {
	// Get id off url
	id := c.Param("id")

	// Get the data off req body
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	// Find the post were updating
	var post models.Post
	result1 := initializers.DB.First(&post, id)

	// Throw error
	if result1.Error != nil {
		response := fmt.Sprintf("the post with id %s does not exists.", id)
		c.JSON(404, gin.H{
			"message": response,
		})
		return
	}

	// Update it
	result2 := initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})
	if result2.Error != nil {
		c.JSON(500, gin.H{
			"message": "An unexpected error occurred while processing the request.",
		})
		return
	}

	// Respond with it
	c.JSON(200, gin.H{
		"posts": post,
	})
}

func DeletePost(c *gin.Context) {
	// Get id off url
	id := c.Param("id")

	// Find the post were updating
	var post models.Post
	result1 := initializers.DB.First(&post, id)

	// Throw error
	if result1.Error != nil {
		response := fmt.Sprintf("the post with id %s does not exists.", id)
		c.JSON(404, gin.H{
			"message": response,
		})
		return
	}

	// Delete the post
	result2 := initializers.DB.Delete(&models.Post{}, id)

	// Throw error
	if result2.Error != nil {
		c.JSON(500, gin.H{
			"message": "An unexpected error occurred while processing the request.",
		})
		return
	}

	// Respond with it
	c.JSON(200, gin.H{
		"post deleted": post,
	})
}
