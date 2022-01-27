package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"avarice/models"
)

type CreateSwordInput struct {
	Name			string		`json:"name" binding:"required"`
	Image			string		`json:"image" binding:"required"`
	Price			float32		`json:"price,string" binding:"required"`
	Inches			int			`json:"inches,string" binding:"required"`
	Ounces			int			`json:"ounces,string" binding:"required"`
	Mats			string		`json:"mats" binding:"required"`
	Description		string		`json:"description" binding:"required"`
}

type UpdateSwordInput struct {
	Name			string		`json:"name"`
	Image			string		`json:"image"`
	Price			float32		`json:"price,string"`
	Inches			int			`json:"inches,string"`
	Ounces			int			`json:"ounces,string"`
	Mats			string		`json:"mats"`
	Description		string		`json:"description"`
}

// GET all swords (Index)
func FindSwords(c *gin.Context) {
	var swords []models.Sword
	models.DB.Find(&swords)

	c.JSON(http.StatusOK, gin.H{"data": swords})
}

// POST sword (Create)
func CreateSword(c *gin.Context) {
	var input CreateSwordInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sword := models.Sword{
		Name: input.Name,
		Image: input.Image,
		Price: input.Price,
		Inches: input.Inches,
		Ounces: input.Ounces,
		Mats: input.Mats,
		Description: input.Description,
	}
	models.DB.Create(&sword)

	c.JSON(http.StatusOK, gin.H{"data": sword})
}

// GET one sword (Show)
func FindSword(c *gin.Context) {
	var sword models.Sword

	if err := models.DB.Where("id = ?", c.Param("id")).First(&sword).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": sword})
}

// PATCH one sword (Update)
func UpdateSword(c *gin.Context) {
	var sword models.Sword
	if err := models.DB.Where("id = ?", c.Param("id")).First(&sword).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	var input UpdateSwordInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&sword).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": sword})

}

// DELETE one sword (Delete)
func DeleteSword(c *gin.Context) {
	var sword models.Sword
	if err := models.DB.Where("id = ?", c.Param("id")).First(&sword).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	models.DB.Delete(&sword)

	c.JSON(http.StatusOK, gin.H{"data": true})
}