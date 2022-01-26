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
	r := gin.New()
	r.Use(CORS)

	swords := r.Group("/swords")
	swords.GET("/", controllers.FindSwords)
	swords.POST("/", controllers.CreateSword)
	swords.GET("/:id", controllers.FindSword)
	swords.PATCH("/:id", controllers.UpdateSword)
	swords.DELETE("/:id", controllers.DeleteSword)
	models.ConnectDatabase()
	r.Run()
}