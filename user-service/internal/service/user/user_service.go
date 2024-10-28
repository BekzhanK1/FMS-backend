package user

import (
	"fmt"
	"log"
	"time"
	"user-service/internal/models"
	"user-service/internal/utils"
	"user-service/types"
)

type Service struct {
	userStore       types.UserStore
	otpStore        types.OTPStore
	farmerInfoStore types.FarmerInfoStore
}

func NewService(userStore types.UserStore, otpStore types.OTPStore, farmerInfoStore types.FarmerInfoStore) *Service {
	return &Service{
		userStore:       userStore,
		otpStore:        otpStore,
		farmerInfoStore: farmerInfoStore,
	}
}

func (h *Service) CreateUser(email, username, first_name, last_name, phone, passwordHash string, isActive bool, role models.Role, profilePicture string) (string, error) {
	if role == models.Admin {
		return "", fmt.Errorf("cannot create admin user")
	}

	userObject := &models.User{
		Email:          email,
		Username:       username,
		FirstName:      first_name,
		LastName:       last_name,
		Phone:          phone,
		PasswordHash:   passwordHash,
		IsActive:       false,
		Role:           role,
		ProfilePicture: profilePicture,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	user, err := h.userStore.CreateUser(userObject)
	if err != nil {
		return "", err
	}

	if user.Role == models.Farmer {
		farmerInfo := &models.FarmerInfo{
			FarmerID:   user.ID,
			Rating:     0.0,
			Experience: 0,
			Bio:        "",
			IsVerified: false,
		}
		err = h.farmerInfoStore.CreateFarmerInfo(farmerInfo)
		if err != nil {
			return "", err
		}
	}

	OtpCode, encryptedEmail, err := h.otpStore.CreateOTP(user)
	if err != nil {
		return "", err
	}
	utils.SendEmail(email, "Your OTP: %s", OtpCode)

	return encryptedEmail, nil
}

func (s *Service) GetUserById(id int) (*models.User, error) {
	user, err := s.userStore.GetUserById(id)
	if err != nil {
		log.Fatalf("error: %s", err)
		log.Fatalf("could not get user with id %d", id)
		return nil, err
	}

	return user, nil
}

func (s *Service) UpdateUser(id int, username, phone, profilePicture string, isActive bool) error {
	user := &models.User{
		Username:       username,
		Phone:          phone,
		IsActive:       isActive,
		ProfilePicture: profilePicture,
		UpdatedAt:      time.Now(),
	}

	if err := s.userStore.UpdateUser(id, user); err != nil {
		log.Fatalf("could not update user with id %d", id)
		return err
	}

	return nil
}

func (s *Service) DeleteUser(id int) error {
	if err := s.userStore.DeleteUser(id); err != nil {
		log.Fatalf("could not delete user with id %d", id)
		return err
	}

	return nil
}

func (s *Service) GetUserByEmail(email string) (*models.User, error) {
	user, err := s.userStore.GetUserByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

func (s *Service) ActivateUser(encryptedEmail string, otpCode string) error {
	userEmail, err := utils.Decrypt(encryptedEmail)
	if err != nil {
		return fmt.Errorf("could not decrypt email: %w", err)
	}

	user, err := s.userStore.GetUserByEmail(userEmail)

	if user.IsActive {
		return fmt.Errorf("user is already activated: %w", err)
	}

	otp, err := s.otpStore.GetOTPByUserId(user.ID)
	if err != nil {
		return fmt.Errorf("could not get token: %w", err)
	}

	if otp.ExpiresAt.Before(utils.GetCurrentTime()) {
		newOtpCode := utils.GenerateOTP()
		err = s.otpStore.RegenerateOTP(user.ID, newOtpCode)
		if err != nil {
			return fmt.Errorf("couldn't regenerate OTP: %w", err)
		}

		utils.SendEmail(user.Email, "OTP Code", newOtpCode)

		return fmt.Errorf("OTP is expired, a new OTP has been sent to email %s", user.Email)
	}

	fmt.Printf("otpCode: %s, otp.OTP_Code: %s\n", otpCode, otp.OTP_Code)

	if otpCode != otp.OTP_Code {
		return fmt.Errorf("invalid OTP code")
	}

	user.IsActive = true
	err = s.userStore.UpdateUser(user.ID, user)
	if err != nil {
		return err
	}

	if err = s.otpStore.DeleteOTP(user.ID); err != nil {
		return err
	}

	return nil
}
