package models

import (
	"time"
)

type PersonalAccount struct {
	Account   Account   `gorm:"embedded"`
	Bio       string    `json:"bio" gorm:"type:text"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	// Posts          []Post    `gorm:"foreignKey:UserID"`
	// Follows        []Follow  `gorm:"foreignKey:FollowerID"`
	// Comments       []Comment `gorm:"foreignKey:UserID"`
	// Likes          []Like    `gorm:"foreignKey:UserID"`
}
