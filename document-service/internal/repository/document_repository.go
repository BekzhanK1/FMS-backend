package repository

import (
	"context"
	"document-service/internal/models"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Store struct {
	db *mongo.Collection
}

func NewStore(db *mongo.Database) *Store {
	return &Store{
		db: db.Collection("documents"),
	}
}

func (s *Store) GetByFarmerID(ctx context.Context, farmerId int) (*models.Document, error) {
	document := &models.Document{}
	if err := s.db.FindOne(ctx, bson.M{"farmer_id": farmerId}).Decode(document); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, err
		}
	}

	return document, nil
}

func (s *Store) CreateDocument(ctx context.Context, data *models.Document) (string, error) {
	res, err := s.db.InsertOne(ctx, data)
	if err != nil {
		return "", nil
	}

	
	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}
