package shelter

import (
	"net/http"
	"pumplepet-server/internal/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetShelters retrieves all shelters from the database
func GetShelters(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var shelters []model.Shelter

	result := db.Preload("Pets").Find(&shelters)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch shelters"})
		return
	}

	c.JSON(http.StatusOK, shelters)
}

// GetShelterById retrieves a specific shelter by ID
func GetShelterById(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	var shelter model.Shelter

	result := db.Preload("Pets").First(&shelter, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Shelter not found"})
		return
	}

	c.JSON(http.StatusOK, shelter)
}
