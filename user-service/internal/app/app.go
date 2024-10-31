package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	db "user-service/internal/database"
	store "user-service/internal/repository"
	applService "user-service/internal/service/application"
	authService "user-service/internal/service/auth"
	farmService "user-service/internal/service/farms"
	userService "user-service/internal/service/user"
	httpHandler "user-service/internal/transport/http"
	adminutils "user-service/internal/utils/adminutils"
)

func Run() {
	db, err := db.Connect()
	if err != nil {
		log.Fatalf("error occurred connecting to db: %s", err)
	}
	defer db.Close()

	r := mux.NewRouter()

	// Setup services and handlers as before
	userStore := store.NewUserStore(db)
	tokenStore := store.NewTokenStore(db)
	otpStore := store.NewOTPStore(db)
	farmerInfoStore := store.NewFarmerInfoStore(db)
	buyerInfoStore := store.NewBuyerInfoStore(db)
	farmStore := store.NewFarmStore(db)
	applicationStore := store.NewApplicationStore(db)

	adminutils.CreateAdminUserIfNotExists(userStore)

	userService := userService.NewService(userStore, otpStore, farmerInfoStore, buyerInfoStore)
	authService := authService.NewService(tokenStore)
	farmService := farmService.NewService(farmStore, userStore, applicationStore)
	applService := applService.NewService(farmStore, userStore, applicationStore)
	userHandler := httpHandler.NewHanlder(*userService, *authService, *farmService, *applService)

	userRouter := r.PathPrefix("/users").Subrouter()
	userHandler.RegisterRoutes(userRouter)

	// Apply CORS middleware to the main router
	corsRouter := enableCors(r)

	log.Println("Starting server on :8080")
	err = http.ListenAndServe(":8080", corsRouter) // Use corsRouter here
	if err != nil {
		log.Fatal(err)
	}
}

func enableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
