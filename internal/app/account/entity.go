package account

import "time"

type Account struct {
	ID             uint   `gorm:"primaryKey"`
	Username       string `gorm:"unique;not null"`
	FirstName      string `gorm:"not null"`
	LastName       string `gorm:"not null"`
	Email          string `gorm:"unique;not null"`
	PasswordHash   string `gorm:"not null"`
	ProfilePicture string `gorm:"size:255"` // URL to their profile picture
	Bio            string `gorm:"type:text"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	// Cafes          []Cafe    `gorm:"foreignKey:OwnerID"` // User can own multiple cafes
	// Posts          []Post    `gorm:"foreignKey:UserID"`
	// Follows        []Follow  `gorm:"foreignKey:FollowerID"`
	// Comments       []Comment `gorm:"foreignKey:UserID"`
	// Likes          []Like    `gorm:"foreignKey:UserID"`
}
