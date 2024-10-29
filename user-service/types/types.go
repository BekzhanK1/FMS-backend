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

type BuyerInfoStore interface {
	CreateBuyerInfo(*models.BuyerInfo) error
	UpdateBuyerInfo(*models.BuyerInfo) error
	GetBuyerInfoByBuyerID(int) (*models.BuyerInfo, error)
	DeleteBuyerInfo(int) error
}

type FarmStore interface {
	CreateFarm(farm *models.Farm) (*models.Farm, error)
	GetFarmByID(id int) (*models.Farm, error)
	UpdateFarm(farmerID int, farm *models.Farm) error
	DeleteFarm(farmerID, id int) error
	ListFarms() ([]*models.Farm, error)
	ListFarmsByFarmerID(farmerID int) ([]*models.Farm, error)
}

type ApplicationStore interface {
	CreateApplication(application *models.Application) error
	ListApplications() ([]*models.Application, error)
	ListApplicationsWithDetails() ([]*ApplicationResponse, error)
	UpdateApplication(id int, status string, rejectionReason *string) error
	GetApplicationByID(id int) (*models.Application, error)
	GetApplicationByFarmID(farmID int) (*models.Application, error)
	GetApplicationsByFarmerID(farmerID int) ([]*models.Application, error)
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

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

type LoginPayload struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type CreateFarmPayload struct {
	Name       		string `json:"name" validate:"required"`
	Address    		string `json:"address"`
	GeoLoc     		string `json:"geo_loc"`
	Size       		string `json:"size"`
	CropTypes  		string `json:"crop_types"`
	// FarmerDocument  *multipart.FileHeader `json:"farmer_document" validate:"required"`
	// FarmDocument    *multipart.FileHeader `json:"farm_documment" validate:"required"`
}





// RESPONSE STRUCTS

type ApplicationResponse struct {
    ID              int                `json:"id"`
    Status          models.ApplicationStatus  `json:"status"`
    RejectionReason string             `json:"rejection_reason"`
    CreatedAt       time.Time          `json:"created_at"`
    Farmer          FarmerResponse     `json:"farmer"`
    Farm            FarmDetails        `json:"farm"`
}

type FarmerResponse struct {
    ID             int     `json:"id"`
    FirstName      string  `json:"first_name"`
    LastName       string  `json:"last_name"`
    Username       string  `json:"username"`
    Email          string  `json:"email"`
    Phone          string  `json:"phone_number"`
    ProfilePicture string  `json:"profile_picture_url"`
    Role           models.Role    `json:"role"`
	Rating     float32 `json:"rating"`
    Experience int     `json:"experience"`
    Bio        string  `json:"bio"`
}

type FarmDetails struct {
    ID        int    `json:"id"`
    Name      string `json:"name"`
    Address   string `json:"address"`
    GeoLoc    string `json:"geo_loc"`
    Size      string `json:"size"`
    CropTypes string `json:"crop_types"`
}

