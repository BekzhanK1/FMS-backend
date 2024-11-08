package adminutils

import (
	"log"
	"time"
	"user-service/internal/config"
	"user-service/internal/models"
	store "user-service/internal/repository"
)

func CreateAdminUserIfNotExists(userStore *store.UserStore) {
	existingUser, err := userStore.GetUserByEmail(config.AdminConfig.Email)
	if err != nil {
		log.Fatalf("could not get user by email: %s", err)
	}

	if existingUser != nil {
		log.Println("admin user already exists")
		log.Printf("Admin email: %s", existingUser.Email)
		return
	}

	adminUser := &models.User{
		Email:          config.AdminConfig.Email,
		Username:       config.AdminConfig.Username,
		Phone:          config.AdminConfig.Phone,
		PasswordHash:   config.AdminConfig.Password,
		IsActive:       true,
		Role:           models.Admin,
		ProfilePicture: "",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	_, err = userStore.CreateUser(adminUser)
	if err != nil {
		log.Fatalf("could not create admin user: %s", err)
	}
	log.Println("admin user created")

}