package types

import (
	"time"
	"user-service/internal/models"
)

type UserStore interface {
	CreateUser(*models.User) error
	DeleteUser(int) error
	GetUserById(int) (*models.User, error)
	UpdateUser(int, *models.User) error
	GetUserByEmail(string) (*models.User, error)
}

type TokenStore interface {
	GetTokenByUserId(int) (*models.Token, error)
	CreateToken(*models.Token) error
	UpdateTokenByUserId(int, *models.Token) error
}

type CreateUserPayload struct {
	Email          string      `json:"email" validate:"required"`
	Username       string      `json:"username" validate:"omitempty"`
	Phone          string      `json:"phone" validate:"required"`
	Password   string      `json:"password" validate:"required"`
	Role           models.Role `json:"role" validate:"required,oneof=Farmer Buyer Admin"`
	ProfilePicture string      `json:"profile_picture" validate:"omitempty"`
}

type UpdateUserPayload struct {
	Username       string `json:"username" validate:"required"`
	Phone          string `json:"phone" validate:"required"`
	ProfilePicture string `json:"profile_picture" validate:"required"`
	IsActive       bool   `json:"is_active" validate:"omit_empty"`
}

type CreateTokenPayload struct {
	UserID     int
	Token      string
	Expiration time.Time
	UpdatedAt  time.Time
}

type UpdateTokenPayload struct {
	Token      string
	Expiration time.Time
	UpdatedAt  time.Time
}

type LoginPayload struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

