package model

import (
	"time"
	"gorm.io/gorm"
)

// Pet represents a pet available for adoption
type Pet struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `gorm:"not null" json:"name"`
	Species     string         `gorm:"not null" json:"species"`
	Breed        string         `json:"breed"`
	Age          float64        `json:"age"` // Age in years
	Gender       string         `json:"gender"`
	Size         string         `json:"size"` // Small, Medium, Large
	Description  string         `gorm:"type:text" json:"description"`
	Status       string         `gorm:"default:available" json:"status"` // available, pending, adopted
	UserID       uint           `json:"user_id,omitempty"` // Foreign key to the owner (if adopted)
	ShelterID     uint           `json:"shelter_id"` // Shelter where the pet is located
	AdoptedAt      *time.Time     `json:"adopted_at,omitempty"` // When the pet was adopted (if at all)
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`

	// Metadata
	PetMetadata PetMetadata `gorm:"foreignKey:PetID;references:ID" json:"metadata"`
	
	// Relationships
	Owner    User    `gorm:"foreignKey:UserID" json:"owner,omitempty"` // The pet's owner
	Shelter   Shelter `gorm:"foreignKey:ShelterID" json:"shelter"` // The shelter where the pet is located
}

// PetMetadata stores additional information about a pet
type PetMetadata struct {
	PetID          uint   `gorm:"primaryKey" json:"pet_id"` // Using PetID as the primary key
	Color          string `json:"color"`
	Weight         float64 `json:"weight"`
	MedicalHistory string `gorm:"type:text" json:"medical_history,omitempty"`
	Vaccinated     bool   `json:"vaccinated"`
	Neutered       bool   `json:"neutered"`
	SpecialNeeds   string `json:"special_needs,omitempty"`
} 