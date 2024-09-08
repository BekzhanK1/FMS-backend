package models

import (
	"time"
)


type Role string

const (
	Farmer Role = "Farmer"
	Buyer  Role = "Buyer"
	Admin  Role = "Admin"
)

type User struct {
	ID              int       `json:"id" db:"id"`
	Email           string    `json:"email" db:"email"`
	Username        string    `json:"username" db:"username"`
	Phone           string    `json:"phone" db:"phone"`
	PasswordHash    string    `json:"-" db:"password"`
	IsActive        bool      `json:"is_active" db:"is_active"`
	Role            Role      `json:"role" db:"role"`
	ProfilePicture  string    `json:"profile_picture" db:"profile_picture"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}