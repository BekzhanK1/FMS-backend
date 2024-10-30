package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Document struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`               // MongoDB document ID
	FileName         string             `bson:"file_name"`                   // Original file name
	UploadDate       time.Time          `bson:"upload_date"`                 // Timestamp of upload
	FileSize         int64              `bson:"file_size"`                   // Size in bytes
	FarmerId         int                `bson:"farmer_id"`                   // ID of the user who uploaded the file
	Status           string             `bson:"status"`                      // e.g., "pending", "approved", "rejected"
	VerificationDate *time.Time         `bson:"verification_date,omitempty"` // When it was verified
}
