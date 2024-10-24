package types

import (
	"time"
	"user-service/internal/models"
)

type UserStore interface {
	CreateUser(*models.User) (*models.User, error)
	DeleteUser(int) error
	GetUserById(int) (*models.User, error)
	UpdateUser(int, *models.User) error
	GetUserByEmail(string) (*models.User, error)
	ActivateUser(string, string) error
}

type TokenStore interface {
	GetTokenByUserId(int) (*models.Token, error)
	CreateToken(*models.Token) error
	UpdateTokenByUserId(int, *models.Token) error
}

type OTPStore interface {
	CreateOTP(*models.User) (string, string, error)
	DeleteOTP(int) (error)
	GetOTPByUserId(int) (*models.OTP, error)
	RegenerateOTP(int, string) (error)
}

type FarmerInfoStore interface {
	CreateFarmerInfo(*models.FarmerInfo) error
	UpdateFarmerInfo(*models.FarmerInfo) error
	GetFarmerInfoByFarmerId(int) (*models.FarmerInfo, error)
	DeleteFarmerInfo(int) error
}
	

type CreateUserPayload struct {
	Email          string      `json:"email" validate:"required"`
	Username       string      `json:"username" validate:"omitempty"`
	FirstName      string      `json:"first_name" validate:"omitempty"`
	LastName       string      `json:"last_name" validate:"omitempty"`
	Phone          string      `json:"phone" validate:"required"`
	Password   	   string      `json:"password" validate:"required"`
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

