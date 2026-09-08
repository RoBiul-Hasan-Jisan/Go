package models

import "time"

// User represents a registered application user.
// The password hash is never serialized to JSON responses.
type User struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Name         string    `gorm:"not null" json:"name"`
	Email        string    `gorm:"uniqueIndex;not null" json:"email"`
	PasswordHash string    `gorm:"not null" json:"-"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Tasks        []Task    `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"-"`
}
