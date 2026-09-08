package models

import "time"

// Task represents a single to-do item that belongs to a user.
type Task struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	UserID      uint      `gorm:"not null;index" json:"user_id"`
	Title       string    `gorm:"not null" json:"title"`
	Description string    `json:"description"`
	Completed   bool      `gorm:"not null;default:false" json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
