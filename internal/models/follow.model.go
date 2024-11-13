package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Follow struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	FollowerUUID uuid.UUID `json:"follower_uuid" gorm:"not null"`
	FollowedUUID uuid.UUID `json:"followed_uuid" gorm:"not null"`
	FollowType   string    `json:"followed_type" gorm:"not null"`
	CreatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
