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
	FirstName      string	 `db:"first_name"`
	LastName       string	 `db:"last_name"`
	Phone          string    `db:"phone_number"`
	PasswordHash   string    `db:"password_hash"`
	IsActive       bool      `db:"is_active"`
	Role           Role      `db:"role"`
	ProfilePicture string    `db:"profile_picture_url"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}
