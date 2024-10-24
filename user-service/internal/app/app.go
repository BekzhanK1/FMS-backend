package app

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"user-service/internal/config"
	db "user-service/internal/database"
	"user-service/internal/models"
	store "user-service/internal/repository"
	userService "user-service/internal/service/user"
	authService "user-service/internal/service/auth"
	httpHandler "user-service/internal/transport/http"
)

func Run() {
	db, err := db.Connect()
	if err != nil {
		log.Fatalf("error occurred connecting to db: %s", err)
	}
	defer db.Close()

	r := mux.NewRouter()

	userStore := store.NewUserStore(db)
	tokenStore := store.NewTokenStore(db)
	otpStore := store.NewOTPStore(db)

	createAdminUserIfNotExists(userStore)

	userService := userService.NewService(userStore, otpStore)
	authService := authService.NewService(tokenStore)
	userHandler := httpHandler.NewHanlder(*userService, *authService)

	userRouter := r.PathPrefix("/users").Subrouter()
	userHandler.RegisterRoutes(userRouter)

	log.Println("Starting server on :8080")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}

// TO-DO: Remove it from app file and put somewhere else
func createAdminUserIfNotExists(userStore *store.UserStore) {
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
