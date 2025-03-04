package model

import (
	"time"
	"gorm.io/gorm"
)

// Pet represents a pet available for adoption
type Pet struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"not null" json:"name"`
	Species   string         `gorm:"not null" json:"species"`
	UserID    uint           `json:"user_id"`             // Foreign key to the owner
	AdoptedAt *time.Time     `json:"adopted_at,omitempty"` // When the pet was adopted (if at all)

	// Metadata
	PetMetadata PetMetadata `gorm:"foreignKey:PetID;references:ID" json:"metadata"`
	
	// Relationships
	Owner User `gorm:"foreignKey:UserID" json:"owner,omitempty"` // The pet's owner
}

// PetMetadata stores additional information about a pet
type PetMetadata struct {
	PetID     uint           `gorm:"primaryKey" json:"pet_id"` // Using PetID as the primary key
	Color     string         `json:"color"`
	Weight    float64        `json:"weight"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
} 