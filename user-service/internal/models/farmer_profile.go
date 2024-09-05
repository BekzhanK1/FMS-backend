package models

import (
	"time"
)

type FarmerProfile struct {
	ID           int       `json:"id" db:"id"`
	FarmID       int       `json:"farm_id" db:"farm_id"`
	ProfilePicture string   `json:"profile_picture" db:"profile_picture"`
	ContactPhone string    `json:"contact_phone" db:"contact_phone"`
	Description  string    `json:"description" db:"description"`
	Verified     bool      `json:"verified" db:"verified"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}
