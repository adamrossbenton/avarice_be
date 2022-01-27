package main

import (
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	
	"avarice/models"
	"avarice/controllers"
	"avarice/svc"
)

// JWT Middleware and Controller source: https://medium.com/wesionary-team/jwt-authentication-in-golang-with-gin-63dbc0816d55

func main() {
	var loginService service.LoginService = service.StaticLoginService()
	var jwtService service.JWTService = service.JWTAuthService()
	var loginController controllers.LoginController = controllers.LoginHandler(loginService, jwtService)

	r := gin.Default()
	r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
        AllowHeaders:     []string{"Origin", "Content-Type"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }))

	r.POST("/login", func(c *gin.Context) {
		token := loginController.Login(c)
		if token != "" {
			c.JSON(http.StatusOK, gin.H{"token": token,})
		} else {
			c.JSON(http.StatusUnauthorized, nil)
		}
	})

	r.GET("/", controllers.FindSwords)
	r.POST("/", controllers.CreateSword)
	r.GET("/:id", controllers.FindSword)
	r.PATCH("/:id", controllers.UpdateSword)
	r.DELETE("/:id", controllers.DeleteSword)
	models.ConnectDatabase()
	r.Run()
}