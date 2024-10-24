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
	userStore  types.UserStore
	otpStore   types.OTPStore
}

func NewService(userStore types.UserStore, otpStore types.OTPStore) *Service {
	return &Service{
		userStore:  userStore,
		otpStore:   otpStore,
	}
}

func (h *Service) CreateUser(email, username, phone, passwordHash string, isActive bool, role models.Role, profilePicture string) (string, error) {
	if role == models.Admin {
		return "", fmt.Errorf("cannot create admin user")
	}

	userObject := &models.User{
		Email:          email,
		Username:       username,
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

	OtpCode, encryptedEmail, err := h.otpStore.CreateOTP(user)
	if err != nil {
		return "", err
	}
	utils.SendEmail(email, "Your OTP: %s", OtpCode)

	return encryptedEmail, nil
}

func (h *Service) GetUserById(id int) (*models.User, error) {
	user, err := h.userStore.GetUserById(id)
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

	if err := h.userStore.UpdateUser(id, user); err != nil {
		log.Fatalf("could not update user with id %d", id)
		return err
	}

	return nil
}

func (h *Service) DeleteUser(id int) error {
	if err := h.userStore.DeleteUser(id); err != nil {
		log.Fatalf("could not delete user with id %d", id)
		return err
	}

	return nil
}

func (h *Service) GetUserByEmail(email string) (*models.User, error) {
	user, err := h.userStore.GetUserByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

func (h *Service) ActivateUser(encryptedEmail, otpCode string) error {
	err := h.userStore.ActivateUser(encryptedEmail, otpCode)
	if err != nil {
		return err
	}
	return nil
}
