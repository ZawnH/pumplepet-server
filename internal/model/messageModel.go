package model

import (
	"time"
)

type Message struct {
	ID          uint      `gorm:"primary_key"`
	Content     string    `gorm:"not null"`
	SenderID    uint      `gorm:"not null"`
	RecipientID uint      `gorm:"not null"`
	Sender      User      `gorm:"foreignkey:SenderID"`
	Recipient   User      `gorm:"foreignkey:RecipientID"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
