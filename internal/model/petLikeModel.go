package model

import (
	"time"

	"gorm.io/gorm"
)

// PetLike represents a user's like/dislike status for a pet
type PetLike struct {
	PetID     uint           `gorm:"primaryKey;autoIncrement:false" json:"pet_id"`
	UserID    uint           `gorm:"primaryKey;autoIncrement:false" json:"user_id"`
	Status    bool           `json:"status"` // true for like, false for dislike
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	Pet  Pet  `gorm:"foreignKey:PetID" json:"pet,omitempty"`
	User User `gorm:"foreignKey:UserID" json:"user,omitempty"`
}
