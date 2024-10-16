package models

import "time"

type Comment struct {
	ID        uint   `gorm:"primaryKey"`
	PostID    uint   `gorm:"not null"` // Post the comment belongs to
	Post      Post   `gorm:"foreignKey:PostID"`
	UserID    uint   `gorm:"not null"` // User who commented
	User      User   `gorm:"foreignKey:UserID"`
	Content   string `gorm:"type:text;not null"`
	CreatedAt time.Time
}
