package account

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	UUID           uuid.UUID `json:"uuid" gorm:"primaryKey;default:uuid_generate_v4()"`
	Username       string    `json:"username" gorm:"unique;not null"`
	FirstName      string    `json:"first_name" gorm:"not null"`
	LastName       string    `json:"last_name" gorm:"not null"`
	Email          string    `json:"email" gorm:"unique;not null"`
	PasswordHash   string    `json:"password_hash" gorm:"not null"`
	ProfilePicture string    `json:"profile_picture" gorm:"size:255"` // URL to their profile picture
	Bio            string    `json:"bio" gorm:"type:text"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	// Cafes          []Cafe    `gorm:"foreignKey:OwnerID"` // User can own multiple cafes
	// Posts          []Post    `gorm:"foreignKey:UserID"`
	// Follows        []Follow  `gorm:"foreignKey:FollowerID"`
	// Comments       []Comment `gorm:"foreignKey:UserID"`
	// Likes          []Like    `gorm:"foreignKey:UserID"`
}
