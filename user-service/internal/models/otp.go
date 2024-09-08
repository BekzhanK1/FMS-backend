package models

import (
	"time"
)

type OTP struct {
	UserID    int    `json:"user_id" db:"user_id"`
	OTP_Code  string `json:"otp_code" db:"otp_code"`
	ExpiresAt time.Time `json:"expires_at" db:"expires_at"`
}