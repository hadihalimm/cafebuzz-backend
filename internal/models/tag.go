package models

import "time"

type Tag struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:50;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Cafes     []Cafe `gorm:"many2many:cafe_tags"` // Associated cafes
}

type CafeTag struct {
	CafeID uint `gorm:"primaryKey"`
	TagID  uint `gorm:"primaryKey"`
}
