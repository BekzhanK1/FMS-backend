package models

type DocumentType string

const (
	NationalID DocumentType = "national_id"
	FarmDocument DocumentType = "farm_document"
)


type Document struct {
	ID 	  int    `json:"id" db:"id"`
	FarmerID int    `json:"farmer_id" db:"farmer_id"`
	Type DocumentType `json:"type" db:"type"`
	Url string `json:"url" db:"url"`
}