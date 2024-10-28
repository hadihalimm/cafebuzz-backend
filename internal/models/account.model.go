package models

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	UUID           uuid.UUID `json:"uuid" gorm:"primaryKey;default:uuid_generate_v4()"`
	Username       string    `json:"username" gorm:"unique;not null"`
	Name           string    `json:"name" gorm:"not null"`
	Email          string    `json:"email" gorm:"unique;not null"`
	PasswordHash   string    `json:"password_hash" gorm:"not null"`
	ProfilePicture string    `json:"profile_picture" gorm:"size:255"` // URL to their profile picture
	Bio            string    `json:"bio" gorm:"type:text"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	// Posts          []Post    `gorm:"foreignKey:UserID"`
	// Follows        []Follow  `gorm:"foreignKey:FollowerID"`
	// Comments       []Comment `gorm:"foreignKey:UserID"`
	// Likes          []Like    `gorm:"foreignKey:UserID"`
}
