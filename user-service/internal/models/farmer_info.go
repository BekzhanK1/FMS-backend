package models

type FarmerInfo struct {
	FarmerID   int  `json:"farmer_id" db:"farmer_id"`
	IsVerified bool `json:"is_verified" db:"is_verified"`
}
