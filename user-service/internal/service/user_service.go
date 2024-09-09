package service

import (
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
		Email:           email,
		Username:        username,
		Phone:           phone,
		PasswordHash:    passwordHash,
		IsActive:        isActive,
		Role:            role,
		ProfilePicture:  profilePicture,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
	err := h.store.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (h *Service) GetUserById(id int) (*models.User, error) {
	return h.store.GetUserById(id)
}

func (h *Service) UpdateUser(id int, email, username, phone, passwordHash string, isActive bool, role models.Role, profilePicture string) error {
	user := &models.User{
		Email:           email,
		Username:        username,
		Phone:           phone,
		PasswordHash:    passwordHash,
		IsActive:        isActive,
		Role:            role,
		ProfilePicture:  profilePicture,
		UpdatedAt:       time.Now(),
	}
	return h.store.UpdateUser(id, user)
}


func (h *Service) DeleteUser(id int) error {
	return h.store.DeleteUser(id)
}
