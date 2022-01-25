package main

import (
	"github.com/gin-gonic/gin"
	
	"avarice/models"
	"avarice/controllers"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/swords", controllers.FindSwords)

	r.Run(":4000")
}