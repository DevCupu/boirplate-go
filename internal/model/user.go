package model

import (
	"time"
)

// User model untuk database
type User struct {
	ID        string     `gorm:"primaryKey" json:"id"`
	Name      string     `json:"name"`
	Email     string     `gorm:"uniqueIndex" json:"email"`
	Phone     string     `json:"phone"`
	Password  string     `json:"-"` // Never expose password
	IsActive  bool       `json:"is_active"`
	LastLogin *time.Time `json:"last_login,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

// TableName sesuaikan nama table
func (User) TableName() string {
	return "users"
}
