package service

import (
	"context"
	"document-service/internal/models"
	"document-service/types"
	"time"

	"log"
)

type Service struct {
	documentStore types.DocumentStore
}

func NewService(documentStore types.DocumentStore) *Service {
	return &Service{
		documentStore,
	}
}

func (s *Service) GetByFarmerID(ctx context.Context, farmerId int) (*models.Document, error) {
	document, err := s.documentStore.GetByFarmerID(ctx, farmerId)
	if err != nil {
		log.Fatalf("error when getting document: %s", err)
		return nil, err
	}
	return document, nil
}

func (s *Service) CreateDocument(ctx context.Context, payload types.CreateDocumentPayload) (string, error) {
	document := &models.Document{
		FileName:   payload.FileName,
		UploadDate: time.Now(),
		FileSize:   0,
		FarmerId:   payload.FarmerId,
		Status:     string(types.Pending),
	}

	id, err := s.documentStore.CreateDocument(ctx, document)
	if err != nil {
		return "", nil
	}

	return id, nil
}
