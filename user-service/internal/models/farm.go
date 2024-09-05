package models

import (
	"time"
)

type Farm struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Address   string    `json:"address" db:"address"`
	Size      string    `json:"size" db:"size"`
	CropTypes string    `json:"crop_types" db:"crop_types"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
