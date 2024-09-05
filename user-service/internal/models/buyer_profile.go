// internal/models/buyer_profile.go

package models

import (
	"time"
)

// BuyerProfile represents the buyer_profile table in the database
type BuyerProfile struct {
	ID             int       `json:"id" db:"id"`
	DeliveryAddress string    `json:"delivery_address" db:"delivery_address"`
	PaymentMethod  string    `json:"payment_method" db:"payment_method"`
	ProfilePicture string    `json:"profile_picture" db:"profile_picture"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}
