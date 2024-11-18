package user

import (
	"fmt"
	"html/template"
	"log"
	"time"
	"user-service/internal/helpers"
	"user-service/internal/models"
	"user-service/internal/templates"
	"user-service/types"
)

type Service struct {
	userStore       types.UserStore
	otpStore        types.OTPStore
	farmerInfoStore types.FarmerInfoStore
	buyerInfoStore  types.BuyerInfoStore
}

func NewService(userStore types.UserStore, otpStore types.OTPStore, farmerInfoStore types.FarmerInfoStore, buyerInfoStore types.BuyerInfoStore) *Service {
	return &Service{
		userStore:       userStore,
		otpStore:        otpStore,
		farmerInfoStore: farmerInfoStore,
		buyerInfoStore:  buyerInfoStore,
	}
}

func (s *Service) CreateUser(email, username, first_name, last_name, phone, passwordHash string, isActive bool, role models.Role, profilePicture string) (string, error) {
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

	user, err := s.userStore.CreateUser(userObject)
	if err != nil {
		return "", err
	}

	if user.Role == models.Farmer || user.Role == models.Buyer {
		farmerInfo := &models.FarmerInfo{
			FarmerID:   user.ID,
			Rating:     0.0,
			Experience: 0,
			Bio:        "",
		}

		buyerInfo := &models.BuyerInfo{
			BuyerID:         user.ID,
			DeliveryAddress: "",
			PaymentMethod:   "",
		}
		err = s.buyerInfoStore.CreateBuyerInfo(buyerInfo)
		if err != nil {
			return "", err
		}

		err = s.farmerInfoStore.CreateFarmerInfo(farmerInfo)
		if err != nil {
			return "", err
		}
	}

	OtpCode, encryptedEmail, err := s.otpStore.CreateOTP(user)
	if err != nil {
		return "", err
	}
	otpData := helpers.OTPData{
		OtpCode: OtpCode,
	}

	t, err := template.New("").Parse(templates.ActivationOtp)
	if err != nil {
		return "", fmt.Errorf("could not load template: %v", err)
	}
	helpers.SendEmail(user.Email, "Your OTP Code", otpData, t)

	return encryptedEmail, nil
}

func (s *Service) GetUserById(id int) (*types.UserResponse, error) {
	user, err := s.userStore.GetUserById(id)
	userResponse := types.UserResponse{
		ID:             user.ID,
		Email:          user.Email,
		Username:       user.Username,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		Phone:          user.Phone,
		ProfilePicture: user.ProfilePicture,
		Role:           user.Role,
		IsActive:       user.IsActive,
	}
	if err != nil {
		log.Fatalf("error: %s", err)
		log.Fatalf("could not get user with id %d", id)
		return nil, err
	}

	return &userResponse, nil
}

func (s *Service) UpdateUser(id int, username, phone, profilePicture string, isActive bool) error {
	user, err := s.userStore.GetUserById(id)

	if err != nil {
		return fmt.Errorf("could not get user: %w", err)
	}

	user.Username = username
	user.Phone = phone
	user.ProfilePicture = profilePicture

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
	userEmail, err := helpers.Decrypt(encryptedEmail)
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

	if otp.ExpiresAt.Before(helpers.GetCurrentTime()) {
		newOtpCode := helpers.GenerateOTP()
		err = s.otpStore.RegenerateOTP(user.ID, newOtpCode)
		if err != nil {
			return fmt.Errorf("couldn't regenerate OTP: %w", err)
		}

		otpData := helpers.OTPData{
			OtpCode: newOtpCode,
		}

		t, err := template.New("").Parse(templates.ActivationOtp)
		if err != nil {
			return fmt.Errorf("could not load template: %v", err)
		}
		helpers.SendEmail(user.Email, "Your OTP Code", otpData, t)

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

func (s *Service) SwitchUserRole(userID int, role string) error {
	roleEnum, err := models.ParseRole(role)
	if err != nil {
		return fmt.Errorf("invalid role")
	}

	if roleEnum == models.Admin {
		return fmt.Errorf("role is not allowed")
	}

	user, err := s.userStore.GetUserById(userID)
	if err != nil {
		return fmt.Errorf("could not retrieve user: %w", err)
	}

	if user == nil {
		return fmt.Errorf("user not found")
	}

	if user.Role == roleEnum || user.Role == models.Admin {
		return fmt.Errorf("user already has the role %s", role)
	}

	user.Role = roleEnum
	if err = s.userStore.UpdateUser(userID, user); err != nil {
		return fmt.Errorf("could not update user role: %w", err)
	}

	return nil
}
