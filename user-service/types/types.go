package types

import "user-service/internal/models"

type UserStore interface {
	CreateUser(*models.User) error
	DeleteUser(int) error
	GetUserById(int) (*models.User, error)
	UpdateUser(int, *models.User) error
}

type CreateUserPayload struct {
	Email          string      `json:"email" validate:"required"`
	Username       string      `json:"username" validate:"omitempty"`
	Phone          string      `json:"phone" validate:"required"`
	Role           models.Role `json:"role" validate:"required"`
	ProfilePicture string      `json:"profile_picture" validate:"omitempty"`
}

type UpdateUserPayload struct {
	Username       string      `json:"username" validate:"required"`
	Phone          string      `json:"phone" validate:"required"`
	ProfilePicture string      `json:"profile_picture" validate:"required"`
}