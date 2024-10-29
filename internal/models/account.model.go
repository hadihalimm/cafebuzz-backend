package models

import "github.com/google/uuid"

type Account struct {
	UUID           uuid.UUID `json:"uuid" gorm:"primaryKey;default:uuid_generate_v4()"`
	Username       string    `json:"username" gorm:"unique;not null"`
	Name           string    `json:"name" gorm:"not null"`
	Email          string    `json:"email" gorm:"unique;not null"`
	PasswordHash   string    `json:"password_hash" gorm:"not null"`
	ProfilePicture string    `json:"profile_picture" gorm:"size:255"` // URL to their profile picture
}
