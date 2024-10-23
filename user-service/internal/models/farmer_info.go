package models

type FarmerInfo struct {
	FarmerID   int  `json:"farmer_id" db:"farmer_id"`
	Rating     float32 `json:"rating" db:"rating"`
	Experience int `json:"experience" db:"experience"`
	Bio 	  string `json:"bio" db:"bio"`
	IsVerified bool `json:"is_verified" db:"is_verified"`
}
