package types

import "user-service/internal/models"


type UserStore interface {
	CreateUser(*models.User) error
	DeleteUser(int) error
	GetUserById(int) (*models.User, error)
	UpdateUser(int, *models.User) error
}