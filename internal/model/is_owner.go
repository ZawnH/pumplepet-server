package model

import (
	"time"
)

// IsOwner represents the relationship between users and pets
// This is a join table for the many-to-many relationship
type IsOwner struct {
	UserID    uint      `gorm:"primaryKey" json:"user_id"`
	PetID     uint      `gorm:"primaryKey" json:"pet_id"`
	AdoptedAt time.Time `json:"adopted_at"`
	
	// Relationships for easier querying
	User User `gorm:"foreignKey:UserID" json:"-"`
	Pet  Pet  `gorm:"foreignKey:PetID" json:"-"`
} 