package main

import (
	"github.com/gin-gonic/gin"

	"go-mongodb/controllers"
)

func main() {
	router := gin.Default()

	router.POST("/mongo", controllers.InsertDB)
	router.GET("/mongo/:id", controllers.SearchById)
	router.GET("/mongo", controllers.SearchAll)
	router.PUT("/mongo", controllers.UpdateDB)
	router.DELETE("/mongo", controllers.DeleteDB)

	router.Run(":9999")
}
