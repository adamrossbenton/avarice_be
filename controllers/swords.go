package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"avarice/models"
)

type CreateSwordInput struct {
	Name			string		`json:"name" binding:"required"`
	Image			string		`json:"image" binding:"required"`
	Price			float32		`json:"price" binding:"required"`
	Inches			int			`json:"Inches" binding:"required"`
	Ounces			int			`json:"ounces" binding:"required"`
	Mats			string		`json:"mats" binding:"required"`
	Description		string		`json:"description" binding:"required"`
}

func FindSwords(c *gin.Context) {
	var swords []models.Sword
	models.DB.Find(&swords)

	c.JSON(http.StatusOK, gin.H{"data": swords})
}

func CreateSword(c *gin.Context) {
	var input CreateSwordInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sword := models.Sword{
		Name: input.Name,
		Price: input.Price,
		Inches: input.Inches,
		Ounces: input.Ounces,
		Mats: input.Mats,
		Description: input.Description,
	}
	models.DB.Create(&sword)

	c.JSON(http.StatusOK, gin.H{"data": sword})
}