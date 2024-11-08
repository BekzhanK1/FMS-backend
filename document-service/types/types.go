package types

import (
	"context"
	"io"
	"mime/multipart"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DocumentStore interface {
	GetFileIDs(context.Context) ([]primitive.ObjectID, error)
	GetFileByID(context.Context, primitive.ObjectID, io.Writer) error
	CreateFile(context.Context, CreateDocumentPayload) (string, error)
}

type CreateDocumentPayload struct {
	FileHeader multipart.FileHeader
	File       multipart.File
}

type Status string

const (
	Pending  Status = "pending"
	Approved Status = "approved"
	Rejected Status = "rejected"
)
