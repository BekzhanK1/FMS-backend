package service

import (
	"context"
	"document-service/types"
	"io"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service struct {
	documentStore types.DocumentStore
}

func NewService(documentStore types.DocumentStore) *Service {
	return &Service{
		documentStore,
	}
}

func (s *Service) GetFileIDs(ctx context.Context) ([]primitive.ObjectID, error) {
	ids, err := s.documentStore.GetFileIDs(ctx)
	if err != nil {
		log.Fatalf("error when getting ids: %s", err)
		return nil, err
	}

	return ids, nil
}

func (s *Service) GetFileByID(ctx context.Context, id primitive.ObjectID, destination io.Writer) error {
	err := s.documentStore.GetFileByID(ctx, id, destination)
	if err != nil {
		log.Fatalf("error when getting file: %s", err)
		return err
	}
	return nil
}

func (s *Service) CreateFile(ctx context.Context, payload types.CreateDocumentPayload) (string, error) {
	id, err := s.documentStore.CreateFile(ctx, payload)
	if err != nil {
		log.Fatalf("error when creating file: %s", err)
		return "", nil
	}

	return id, nil
}
