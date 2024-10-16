package models

import "time"

type Like struct {
	ID        uint `gorm:"primaryKey"`
	PostID    uint `gorm:"not null"` // Post being liked
	Post      Post `gorm:"foreignKey:PostID"`
	UserID    uint `gorm:"not null"` // User who liked the post
	User      User `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
}
