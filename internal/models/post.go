package models

import "time"

type Post struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"not null"` // User who created the post
	User      User   `gorm:"foreignKey:UserID"`
	CafeID    uint   `gorm:"not null"` // Cafe the post is about
	Cafe      Cafe   `gorm:"foreignKey:CafeID"`
	Caption   string `gorm:"type:text"`
	ImageURL  string `gorm:"size:255"` // URL to the image
	CreatedAt time.Time
	UpdatedAt time.Time
	Comments  []Comment `gorm:"foreignKey:PostID"`
	Likes     []Like    `gorm:"foreignKey:PostID"`
}
