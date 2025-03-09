package user

import (
	"net/http"
	"pumplepet-server/internal/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetUsers retrieves all users from the database
func GetUsers(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var users []model.User

	result := db.Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	c.JSON(http.StatusOK, users)
}

// GetUserById retrieves a specific user by ID
func GetUserById(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	var user model.User

	result := db.First(&user, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
