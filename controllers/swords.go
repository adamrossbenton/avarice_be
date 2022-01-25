package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"avarice/models"
)

func FindSwords(c *gin.Context) {
	var swords []models.Sword
	models.DB.Find(&swords)

	c.JSON(http.StatusOK, gin.H{"data": swords})
}