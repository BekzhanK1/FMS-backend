package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	db "user-service/internal/database"
	userStore "user-service/internal/repository"
	"user-service/internal/service"
	"user-service/internal/transport/handler"
)

func Run() {
	db, err := db.Connect()
	if err != nil {
		log.Fatalf("error occurred connecting to db: %s", err)
	}
	defer db.Close()

	r := mux.NewRouter()

	userStore := userStore.NewStore(db)
	userService := service.NewService(userStore)
	userHandler := handler.NewHanlder(*userService)

	userRouter := r.PathPrefix("/users").Subrouter()
	userHandler.RegisterRoutes(userRouter)

	log.Println("Starting server on :8080")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
