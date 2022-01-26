package main

import (

	"github.com/gin-gonic/gin"
	// "github.com/gin-contrib/cors"
	
	"avarice/models"
	"avarice/controllers"
)

// func CORS() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
// 		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
// 		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PATCH, DELETE")
// 		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

// 		if c.Request.Method == "OPTIONS" {
// 			c.AbortWithStatus(204)
// 			return
// 		}

// 		c.Next()
// 	}
// }

func CORS(c *gin.Context) {
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

func main() {
	r := gin.Default()
	r.Use(CORS)
	r.Run()

	r.GET("/swords", controllers.FindSwords)
	r.POST("/swords", controllers.CreateSword)
	r.GET("/swords/:id", controllers.FindSword)
	r.PATCH("/swords/:id", controllers.UpdateSword)
	r.DELETE("/swords/:id", controllers.DeleteSword)

	models.ConnectDatabase()
}