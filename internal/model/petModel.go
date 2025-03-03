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
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Metadata
	PetMetadata PetMetadata `gorm:"foreignKey:PetID" json:"metadata"`
	
	// Relationships
	Owners []User `gorm:"many2many:is_owner" json:"owners,omitempty"`
}

// PetMetadata stores additional information about a pet
type PetMetadata struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	PetID     uint           `gorm:"not null;uniqueIndex" json:"pet_id"`
	Color     string         `json:"color"`
	Weight    float64        `json:"weight"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
} 