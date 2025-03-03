package model

import (
	"time"
	"gorm.io/gorm"
)

// User represents a person looking to adopt a pet
type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Username  string         `gorm:"not null;uniqueIndex" json:"username"`
	Password  string         `gorm:"not null" json:"-"` // Password is not exposed in JSON
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Metadata
	UserMetadata UserMetadata `gorm:"foreignKey:UserID" json:"metadata"`
	
	// Relationships
	Pets []Pet `gorm:"many2many:is_owner" json:"pets,omitempty"`
}

// UserMetadata stores additional information about a user
type UserMetadata struct {
	ID              uint           `gorm:"primaryKey" json:"id"`
	UserID          uint           `gorm:"not null;uniqueIndex" json:"user_id"`
	HomeEnvironment string         `json:"home_environment"`
	LifeStyle       string         `json:"life_style"`
	Preferences     string         `json:"preferences"` // Could also be a JSON field
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
} 