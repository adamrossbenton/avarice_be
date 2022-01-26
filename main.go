package main

import (

	"github.com/gin-gonic/gin"
	
	"avarice/models"
	"avarice/controllers"
)

// Inspiration Source: https://github.com/gin-contrib/cors/issues/29#issuecomment-397859488
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PATCH, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}


func main() {
	r := gin.Default()

	r.GET("/", CORS(), controllers.FindSwords)
	r.POST("/", CORS(), controllers.CreateSword)
	r.GET("/:id", CORS(), controllers.FindSword)
	r.PATCH("/:id", CORS(), controllers.UpdateSword)
	r.DELETE("/:id", CORS(), controllers.DeleteSword)
	models.ConnectDatabase()
	r.Run()
}