package farms

import (
	"context"
	"fmt"
	"time"
	"user-service/internal/models"
	"user-service/internal/service/auth"
	"user-service/types"
)

type FarmService struct {
	farmStore types.FarmStore
	userStore types.UserStore
	applicationStore types.ApplicationStore
}

func NewService(farmStore types.FarmStore, userStore types.UserStore, applicationStore types.ApplicationStore) *FarmService {
	return &FarmService{
		farmStore,
		userStore,
		applicationStore,
	}
}

func (s *FarmService) CreateFarm(ctx context.Context, name, address, geoLoc, size, cropTypes string, isVerified bool) error {
	farmerID, err := auth.GetUserIDFromContext(ctx)
	if err != nil {
		return fmt.Errorf("error retrieving user ID from context: %w", err)
	}
	user, err := s.userStore.GetUserById(farmerID)
	if err != nil {
		return fmt.Errorf("error retrieving user with ID %d: %w", farmerID, err)
	}

	if user == nil {
		return fmt.Errorf("user with ID %d not found", farmerID)
	}

	fmt.Printf("User Role: %v\n", user.Role)
	if user.Role != models.Farmer {
		return fmt.Errorf("user with ID %d is not a farmer", farmerID)
	}

	farm := &models.Farm{
		Name:       name,
		Address:    address,
		GeoLoc:     geoLoc,
		Size:       size,
		CropTypes:  cropTypes,
		IsVerified: isVerified,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		FarmerID:   farmerID,
	}

	farm, err = s.farmStore.CreateFarm(farm)
	if err != nil {
		return fmt.Errorf("could not create farm: %w", err)
	}

	application := &models.Application{
		FarmerID: user.ID,
		FarmID: farm.ID,
	}

	err = s.applicationStore.CreateApplication(application)
	if err != nil {
		return fmt.Errorf("could not create application: %w", err)
	}

	return nil
}


func (s *FarmService) GetFarmByID(id int) (*models.Farm, error) {
	farm, err := s.farmStore.GetFarmByID(id)
	if err != nil {
		return nil, err
	}

	return farm, nil
}

func (s *FarmService) ListFarms() ([]*models.Farm, error) {
	farms, err := s.farmStore.ListFarms()
	if err != nil {
		return nil, err
	}

	return farms, nil
}

func (s *FarmService) ListFarmsByFarmerID(farmerID int) ([]*models.Farm, error) {
	farms, err := s.farmStore.ListFarmsByFarmerID(farmerID)
	if err != nil {
		return nil, err
	}

	return farms, nil
}

func (s *FarmService) UpdateFarm(farmerID int, id int, name, address, geoLoc string, size string, cropTypes string, isVerified bool) error {
	farm := &models.Farm{
		ID:        id,
		Name:      name,
		Address:   address,
		GeoLoc:    geoLoc,
		Size:      size,
		CropTypes: cropTypes,
		IsVerified: isVerified,
		UpdatedAt: time.Now(),
	}

	err := s.farmStore.UpdateFarm(farmerID, farm)
	if err != nil {
		return err
	}

	return nil
}

func (s *FarmService) DeleteFarm(farmerID int, id int) error {
	err := s.farmStore.DeleteFarm(farmerID, id)
	if err != nil {
		return err
	}

	return nil
}