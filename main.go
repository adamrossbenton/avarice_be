package main

import (
	"time"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	
	"avarice/models"
	"avarice/controllers"
)

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
        AllowHeaders:     []string{"Origin", "Content-Type"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }))

	r.GET("/", controllers.FindSwords)
	r.POST("/", controllers.CreateSword)
	r.GET("/:id", controllers.FindSword)
	r.PATCH("/:id", controllers.UpdateSword)
	r.DELETE("/:id", controllers.DeleteSword)
	models.ConnectDatabase()
	r.Run()
}