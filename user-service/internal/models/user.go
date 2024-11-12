package models

import (
	"fmt"
	"time"
)

type Role string

const (
	Farmer Role = "FARMER"
	Buyer  Role = "BUYER"
	Admin  Role = "ADMIN"
)

type User struct {
	ID             int       `db:"id"`
	Email          string    `db:"email"`
	Username       string    `db:"username"`
	FirstName      string    `db:"first_name"`
	LastName       string    `db:"last_name"`
	Phone          string    `db:"phone_number"`
	PasswordHash   string    `db:"password_hash"`
	IsActive       bool      `db:"is_active"`
	Role           Role      `db:"role"`
	ProfilePicture string    `db:"profile_picture_url"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}

func ParseRole(role string) (Role, error) {
	switch role {
	case "Farmer":
		return Farmer, nil
	case "Buyer":
		return Buyer, nil
	case "Admin":
		return Admin, nil
	default:
		return "", fmt.Errorf("invalid role: %s", role)
	}
}
