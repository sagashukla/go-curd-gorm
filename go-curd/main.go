package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sagarshukla785/go-curd/initializers"
	"github.com/sagarshukla785/go-curd/controllers"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToMyDB()
}

func main() {
	r := gin.Default()
	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts/:id", controllers.PostShow)
	r.PUT("/posts/:id", controllers.PostUpdate)
	r.DELETE("/posts/:id", controllers.PostDelete)
	r.Run()
}