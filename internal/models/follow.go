package models

import "time"

type Follow struct {
	ID             uint `gorm:"primaryKey"`
	FollowerID     uint `gorm:"not null"` // User who follows
	Follower       User `gorm:"foreignKey:FollowerID"`
	FollowedUserID uint `gorm:"null"` // If following a user
	FollowedUser   User `gorm:"foreignKey:FollowedUserID"`
	FollowedCafeID uint `gorm:"null"` // If following a cafe
	FollowedCafe   Cafe `gorm:"foreignKey:FollowedCafeID"`
	CreatedAt      time.Time
}
