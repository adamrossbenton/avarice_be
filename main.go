package main

import (
	"github.com/gin-gonic/gin"
	
	"avarice/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.Run(":4000")
}