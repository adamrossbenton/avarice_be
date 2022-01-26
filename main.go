package main

import (
	// "fmt"

	"github.com/gin-gonic/gin"
	// "github.com/gin-contrib/cors"
	
	"avarice/models"
	"avarice/controllers"
)

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
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

// func CORS(c *gin.Context) {
// 		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
// 		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
// 		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PATCH, DELETE")
// 		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

// 		if c.Request.Method == "OPTIONS" {
// 			c.AbortWithStatus(200)
// 			return
// 		} else {
// 			c.Next()
// 		}
// }

type testHeader struct {
	AccessControlAllowOrigin		string		`header:"Access-Control-Allow-Origin"`
}

func main() {
	r := gin.Default()

	swords := r.Group("/swords")

	// From documentation, gonna try something real quick
	// swords.GET("/", controllers.FindSwords, func (c *gin.Context) {
	// 	h := testHeader{AccessControlAllowOrigin:"*"}

	// 	if err := c.ShouldBindHeader(&h); err != nil {
	// 		c.JSON(200, err)
	// 	}

	// 	fmt.Printf("%#v\n", h)
	// 	c.JSON(200, gin.H{"Access-Control-Allow-Origin": h.AccessControlAllowOrigin})
	// })

	swords.GET("/", CORS(), controllers.FindSwords)

	models.ConnectDatabase()
	r.Run()
	// swords.POST("/", controllers.CreateSword)
	// swords.GET("/:id", controllers.FindSword)
	// swords.PATCH("/:id", controllers.UpdateSword)
	// swords.DELETE("/:id", controllers.DeleteSword)
	// models.ConnectDatabase()
	// r.Run()
}