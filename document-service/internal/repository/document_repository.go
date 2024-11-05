package repository

import (
	"context"
	"document-service/types"
	"fmt"
	"io"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Store struct {
	db *mongo.Collection
}

func NewStore(db *mongo.Database) *Store {
	return &Store{
		db: db.Collection("fs.files"),
	}
}

func (s *Store) GetFileIDs(ctx context.Context) ([]primitive.ObjectID, error) {
	var fileIDs []primitive.ObjectID

	cursor, err := s.db.Find(ctx, bson.M{}, options.Find().SetProjection(bson.M{"_id": 1}))
	if err != nil {
		return nil, fmt.Errorf("failed to find files in GridFS: %w", err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var result struct {
			ID primitive.ObjectID `bson:"_id"`
		}
		if err := cursor.Decode(&result); err != nil {
			return nil, fmt.Errorf("failed to decode file ID: %w", err)
		}
		fileIDs = append(fileIDs, result.ID)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor encountered an error: %w", err)
	}

	return fileIDs, nil
}

func (s *Store) GetFileByID(ctx context.Context, fileID primitive.ObjectID, destination io.Writer) error {
	bucket, err := gridfs.NewBucket(s.db.Database())
	if err != nil {
		return fmt.Errorf("failed to create GridFS bucket: %w", err)
	}

	downloadStream, err := bucket.OpenDownloadStream(fileID)
	if err != nil {
		return fmt.Errorf("failed to open download stream: %w", err)
	}
	defer downloadStream.Close()

	if _, err := io.Copy(destination, downloadStream); err != nil {
		return fmt.Errorf("failed to write file to destination: %w", err)
	}

	return nil
}

func (s *Store) CreateFile(ctx context.Context, data types.CreateDocumentPayload) (string, error) {
	bucket, err := gridfs.NewBucket(s.db.Database())
	if err != nil {
		return primitive.NilObjectID.Hex(), fmt.Errorf("failed to create GridFS bucket: %w", err)
	}

	uploadStream, err := bucket.OpenUploadStream(data.FileHeader.Filename)
	if err != nil {
		return primitive.NilObjectID.Hex(), fmt.Errorf("failed to open upload stream: %w", err)
	}
	defer uploadStream.Close()

	if _, err := io.Copy(uploadStream, data.File); err != nil {
		return primitive.NilObjectID.Hex(), fmt.Errorf("failed to upload file to GridFS: %w", err)
	}

	fileID := uploadStream.FileID.(primitive.ObjectID).Hex()
	return fileID, nil
}
