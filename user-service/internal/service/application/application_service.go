package application

import (
	"context"
	"fmt"
	"user-service/internal/models"
	"user-service/internal/service/auth"
	"user-service/types"
)

type ApplicationService struct {
	farmStore        types.FarmStore
	userStore        types.UserStore
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

	var isAdmin bool
	if isAdmin, err = s.isAdmin(userId); err != nil {
		return nil, err
	}

	applications, err := s.applicationStore.ListApplications()
	if err != nil {
		return nil, fmt.Errorf("error listing applications: %w", err)
	}
	if isAdmin {
		return applications, nil
	} else {
		var userApplications []*types.ApplicationResponse
		for _, application := range applications {
			if application.Farmer.ID == userId {
				userApplications = append(userApplications, application)
			}
		}
		return userApplications, nil
	}

}

func (s *ApplicationService) GetApplicationByID(ctx context.Context, id int) (*types.ApplicationResponse, error) {
	userId, err := auth.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("error retrieving user ID from context: %w", err)
	}

	application, err := s.applicationStore.GetApplicationByID(id)
	if err != nil {
		return nil, fmt.Errorf("error getting application with ID %d: %w", id, err)
	}
	var isAdmin bool
	if isAdmin, err = s.isAdmin(userId); err != nil {
		return nil, err
	}
	if !isAdmin && application.Farmer.ID != userId {
		return nil, fmt.Errorf("application with ID %d not found", id)
	}
	if application == nil {
		return nil, fmt.Errorf("application not found")
	}

	return application, nil
}

func (s *ApplicationService) ListApplicationsByFarmerID(ctx context.Context, farmerID int) ([]*types.ApplicationResponse, error) {
	userId, err := auth.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("error retrieving user ID from context: %w", err)
	}

	conditions, err := s.isAdmin(userId)
	if err != nil {
		return nil, err
	} else if !conditions && userId != farmerID {
		return nil, fmt.Errorf("user with ID %d is not authorized to list applications of farmer with ID %d", userId, farmerID)
	}

	applications, err := s.applicationStore.ListApplicationsByFarmerID(farmerID)
	if err != nil {
		return nil, fmt.Errorf("error listing applications of farmer with ID %d: %w", userId, err)
	}

	return applications, nil
}

func (s *ApplicationService) UpdateApplication(ctx context.Context, id int, payload types.ApplicationUpdatePayload) error {
	userId, err := auth.GetUserIDFromContext(ctx)
	if err != nil {
		return fmt.Errorf("error retrieving user ID from context: %w", err)
	}

	isAdmin, err := s.isAdmin(userId)
	if err != nil {
		return err
	}
	if !isAdmin {
		return fmt.Errorf("user with ID %d is not authorized to update application with ID %d", userId, id)
	}

	if payload.Status == "" {
		return fmt.Errorf("status is required")
	}

	if payload.Status != "approved" && payload.Status != "rejected" {
		return fmt.Errorf("status must be 'approved', 'rejected'")
	}

	if payload.Status == "approved" {
		payload.RejectionReason = ""
	}

	if payload.Status == "rejected" && payload.RejectionReason == "" {
		return fmt.Errorf("rejection reason is required for rejected applications")
	}

	err = s.applicationStore.UpdateApplication(id, payload.Status, payload.RejectionReason)
	if err != nil {
		return fmt.Errorf("error updating application with ID %d: %w", id, err)
	}
	return nil
}

func (s *ApplicationService) isAdmin(userId int) (bool, error) {
	user, err := s.userStore.GetUserById(userId)
	if err != nil {
		return false, fmt.Errorf("error retrieving user with ID %d: %w", userId, err)
	}
	if user == nil {
		return false, fmt.Errorf("user with ID %d not found", userId)
	}
	if user.Role != models.Admin {
		return false, nil
	}
	return true, nil
}
