package service

import (
	"log"
	"time"
	"user-service/internal/models"
	"user-service/types"
)

type Service struct {
	store types.UserStore
}

func NewService(store types.UserStore) *Service {
	return &Service{
		store,
	}
}

func (h *Service) CreateUser(email, username, phone, passwordHash string, isActive bool, role models.Role, profilePicture string) (*models.User, error) {
	user := &models.User{
		Email:          email,
		Username:       username,
		Phone:          phone,
		PasswordHash:   passwordHash,
		IsActive:       isActive,
		Role:           role,
		ProfilePicture: profilePicture,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	if err := h.store.CreateUser(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (h *Service) GetUserById(id int) (*models.User, error) {
	user, err := h.store.GetUserById(id)
	if err != nil {
		log.Fatalf("could not get user with id %d", id)
		return nil, err
	}

	return user, nil
}

func (h *Service) UpdateUser(id int, username, phone, profilePicture string, isActive bool) error {
	user := &models.User{
		Username:       username,
		Phone:          phone,
		IsActive:       isActive,
		ProfilePicture: profilePicture,
		UpdatedAt:      time.Now(),
	}

	if err := h.store.UpdateUser(id, user); err != nil {
		log.Fatalf("could not update user with id %d", id)
		return err
	}

	return nil
}

func (h *Service) DeleteUser(id int) error {
	if err := h.store.DeleteUser(id); err != nil {
		log.Fatalf("could not delete user with id %d", id)
		return err
	}

	return nil
}
