package main

import (
	"rest-api-assignment/controllers"
	"rest-api-assignment/initializers"

	"github.com/gin-gonic/gin"
)

func init() {

	initializers.LoadEnvVariables()
	initializers.ConnectToDatabse()

}

func main() {

	router := gin.Default()

	router.POST("/posts", controllers.PostsCreate)
	router.GET("/posts", controllers.PostsIndex)
	router.GET("/posts/:id", controllers.PostsShow)
	router.PUT("/posts/:id", controllers.PostsUpdate)
	router.DELETE("/posts/:id", controllers.PostDelete)

	router.Run()

}

//This is a simple implementation of a RESTful API without utilizing any specific architecture. 
//Although it lacks a defined architectural pattern, it serves as a starting point for building APIs. 
//If you would like me to develop the API using a hexagonal architecture, please let me know, and I will gladly update the code accordingly.
