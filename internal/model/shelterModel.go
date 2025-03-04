package model

import (
	"time"
	"gorm.io/gorm"
)

// Shelter represents a pet shelter or rescue organization
type Shelter struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Address   string         `json:"address"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Additional fields not in diagram but might be useful
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	
	// Relationships (not shown in diagram but implied)
	Pets []Pet `json:"pets,omitempty"`
} 