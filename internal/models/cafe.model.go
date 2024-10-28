package models

import (
	"time"

	"github.com/google/uuid"
)

type Cafe struct {
	UUID           uuid.UUID `json:"uuid" gorm:"primaryKey"`
	Username       string    `json:"username" gorm:"unique;not null"`
	Name           string    `json:"name" gorm:"not null"`
	Email          string    `json:"email" gorm:"unique;not null"`
	PasswordHash   string    `json:"password_hash" gorm:"not null"`
	Description    string    `json:"description" gorm:"type:text"`
	Address        string    `json:"address" gorm:"not null"`
	ProfilePicture string    `json:"profile_picture" gorm:"size:255"`  // URL to their profile picture
	IsVerified     bool      `json:"is_verified" gorm:"default:false"` // For official cafes
	CreatedAt      time.Time
	UpdatedAt      time.Time
	// Posts       []Post `gorm:"foreignKey:CafeID"`
	// Tags        []Tag  `gorm:"many2many:cafe_tags"` // Many-to-many relationship with tags
}
