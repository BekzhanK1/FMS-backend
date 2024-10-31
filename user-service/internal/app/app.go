package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	db "user-service/internal/database"
	store "user-service/internal/repository"
	authService "user-service/internal/service/auth"
	userService "user-service/internal/service/user"
	httpHandler "user-service/internal/transport/http"
	"user-service/internal/helpers/adminutils"
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
	farmerInfoStore := store.NewFarmerInfoStore(db)

	adminutils.CreateAdminUserIfNotExists(userStore)

	userService := userService.NewService(userStore, otpStore, farmerInfoStore)
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


