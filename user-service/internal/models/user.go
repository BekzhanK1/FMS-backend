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
	ID             int       `db:"id"`
	Email          string    `db:"email"`
	Username       string    `db:"username"`
	Phone          string    `db:"phone"`
	PasswordHash   string    `db:"password"`
	IsActive       bool      `db:"is_active"`
	Role           Role      `db:"role"`
	ProfilePicture string    `db:"profile_picture"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}
