package models

import (
	"time"
)

type Tabler interface {
	TableName() string
}

func (CafeAccount) TableName() string {
	return "cafes"
}

type CafeAccount struct {
	Account     Account `gorm:"embedded"`
	Description string  `json:"description" gorm:"type:text"`
	Address     string  `json:"address" gorm:"not null"`
	IsVerified  bool    `json:"is_verified" gorm:"default:false"` // For official cafes
	CreatedAt   time.Time
	UpdatedAt   time.Time
	// Posts       []Post `gorm:"foreignKey:CafeID"`
	// Tags        []Tag  `gorm:"many2many:cafe_tags"` // Many-to-many relationship with tags
}
