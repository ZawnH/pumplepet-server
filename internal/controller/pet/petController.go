package pet

import (
	"net/http"
	"pumplepet-server/internal/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetPets retrieves all pets from the database
func GetPets(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var pets []model.Pet

	result := db.Preload("PetMetadata").Preload("Owner").Find(&pets)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch pets"})
		return
	}

	c.JSON(http.StatusOK, pets)
}

// GetPetById retrieves a specific pet by ID
func GetPetById(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	var pet model.Pet

	result := db.Preload("PetMetadata").Preload("Owner").First(&pet, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pet not found"})
		return
	}

	c.JSON(http.StatusOK, pet)
}

// LikePet 
func LikePet(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	var pet model.Pet
	if err := db.First(&pet, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pet not found"})
		return
	}

	// Create or update like status
	like := model.PetLike{
		PetID:  pet.ID,
		Status: true,
	}

	if err := db.Save(&like).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save like"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pet liked successfully"})
}

// DislikePet
func DislikePet(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	var pet model.Pet
	if err := db.First(&pet, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pet not found"})
		return
	}

	// Create or update dislike status
	like := model.PetLike{
		PetID:  pet.ID,
		Status: false,
	}

	if err := db.Save(&like).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save dislike"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pet disliked successfully"})
}
