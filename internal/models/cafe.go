package models

import "time"

type Cafe struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	Location    string `gorm:"not null"`
	Description string `gorm:"type:text"`
	OwnerID     uint   `gorm:"not null"`           // User who owns the cafe
	Owner       User   `gorm:"foreignKey:OwnerID"` // Relation to the owner user
	IsVerified  bool   `gorm:"default:false"`      // For official cafes
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Posts       []Post `gorm:"foreignKey:CafeID"`
	Tags        []Tag  `gorm:"many2many:cafe_tags"` // Many-to-many relationship with tags
}
