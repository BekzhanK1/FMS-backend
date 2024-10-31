package models

import "time"

type ApplicationStatus string

const (
	StatusPending  ApplicationStatus = "pending"
	StatusApproved ApplicationStatus = "approved"
	StatusRejected ApplicationStatus = "rejected"
)

type Application struct {
	ID              int               `json:"id" db:"id"`
	FarmerID        int               `json:"farmer_id" db:"farmer_id"`
	FarmID          int               `json:"farm" db:"farm"`
	Status          ApplicationStatus `json:"status" db:"status"`
	RejectionReason string            `json:"rejection_reason" db:"rejection_reason"`
	CreatedAt       time.Time         `json:"created_at" db:"created_at"`
}
