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
	r.POST("/swords", controllers.CreateSword)
	r.GET("/swords/:id", controllers.FindSword)
	r.PATCH("/swords/:id", controllers.UpdateSword)

	r.Run(":4000")
}