package types

import (
	"context"
	"document-service/internal/models"
)

type DocumentStore interface {
	GetByFarmerID(context.Context, int) (*models.Document, error)
	CreateDocument(context.Context, *models.Document) (string, error)
}

type CreateDocumentPayload struct {
	FileName string `json:"file_name" validate:"required"` // Original file name
	FarmerId int    `json:"farmer_id" validate:"required"` // ID of the user who uploaded the file
}

type Status string

const (
	Pending  Status = "pending"
	Approved Status = "approved"
	Rejected Status = "rejected"
)
