package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sagarshukla785/go-curd/models"
	"github.com/sagarshukla785/go-curd/initializers"
)

func PostsCreate (c *gin.Context) {
	var body struct {
		Body string
		Title string
	}

	c.Bind(&body)

	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DBM.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}
	
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostShow(c *gin.Context){
	id := c.Param("id")
	var post models.Post 
	initializers.DBM.First(&post, id)

	c.JSON(200, gin.H{
		"post":  post,
	})
}

func PostUpdate(c *gin.Context){
	id := c.Param("id")

	var body struct {
		Body string
		Title string
	}

	c.Bind(&body) 

	var post models.Post
	initializers.DBM.First(&post, id)

	initializers.DBM.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body: body.Body,
	})

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostDelete(c *gin.Context){
	id := c.Param("id")

	initializers.DBM.Delete(&models.Post{}, id)

	c.Status(200)
}