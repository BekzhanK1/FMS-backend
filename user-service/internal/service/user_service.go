// internal/service/user_service.go

package service

import (
	"user-service/internal/models"
	"user-service/internal/repository"
	"time"
)

// CreateUser creates a new user
func CreateUser(email, username, phone, passwordHash string, isActive bool, role models.Role, profilePicture string) (*models.User, error) {
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
	err := repository.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetUser retrieves a user by ID
func GetUser(id int) (*models.User, error) {
	return repository.GetUser(id)
}

// UpdateUser updates user information
func UpdateUser(id int, email, username, phone, passwordHash string, isActive bool, role models.Role, profilePicture string) error {
	user := &models.User{
		ID:              id,
		Email:           email,
		Username:        username,
		Phone:           phone,
		PasswordHash:    passwordHash,
		IsActive:        isActive,
		Role:            role,
		ProfilePicture:  profilePicture,
		UpdatedAt:       time.Now(),
	}
	return repository.UpdateUser(user)
}

// DeleteUser removes a user by ID
func DeleteUser(id int) error {
	return repository.DeleteUser(id)
}
