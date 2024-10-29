package application

import (
	"context"
	"fmt"
	"user-service/internal/models"
	"user-service/internal/service/auth"
	"user-service/types"
)

type ApplicationService struct {
	farmStore types.FarmStore
	userStore types.UserStore
	applicationStore types.ApplicationStore
}

func NewService(farmStore types.FarmStore, userStore types.UserStore, applicationStore types.ApplicationStore) *ApplicationService {
	return &ApplicationService{
		farmStore,
		userStore,
		applicationStore,
	}
}

func (s *ApplicationService) ListApplications(ctx context.Context) ([]*types.ApplicationResponse, error) {
	userId, err := auth.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("error retrieving user ID from context: %w", err)
	}

	if err = s.verifyAdmin(userId); err != nil {
		return nil, err
	}

	applications, err := s.applicationStore.ListApplicationsWithDetails()
	if err != nil {
		return nil, fmt.Errorf("error listing applications: %w", err)
	}

	return applications, nil
}

func (s *ApplicationService) verifyAdmin(userId int) error {
	user, err := s.userStore.GetUserById(userId)
	if err != nil {
		return fmt.Errorf("error retrieving user with ID %d: %w", userId, err)
	}
	if user == nil {
		return fmt.Errorf("user with ID %d not found", userId)
	}
	if user.Role != models.Admin {
		return fmt.Errorf("user with ID %d is not an admin", userId)
	}
	return nil
}
