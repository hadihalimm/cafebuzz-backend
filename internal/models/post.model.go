package models

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID          uint64    `json:"id" gorm:"primaryKey"`
	CreatorUUID uuid.UUID `json:"created_by" gorm:"index;not null"`
	CreatorType string    `json:"owner_type" gorm:"not null"`
	ImageURL    string    `json:"image_url" gorm:"size:255;not null"`
	Caption     string    `json:"caption" gorm:"size:500"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
