package model

import (
	"time"
	"gorm.io/gorm"
)

// Shelter represents a pet shelter or rescue organization
type Shelter struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `gorm:"not null" json:"name"`
	Address     string         `json:"address"`
	City        string         `json:"city"`
	State       string         `json:"state"`
	ZipCode     string         `json:"zip_code"`
	PhoneNumber string         `json:"phone_number"`
	Email       string         `gorm:"uniqueIndex" json:"email"`
	Website     string         `json:"website,omitempty"`
	Description string         `gorm:"type:text" json:"description,omitempty"`
	Latitude    float64        `json:"latitude,omitempty"` // For map integration
	Longitude   float64        `json:"longitude,omitempty"` // For map integration
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	
	// Relationships
	Pets []Pet `gorm:"foreignKey:ShelterID" json:"pets,omitempty"`
} 