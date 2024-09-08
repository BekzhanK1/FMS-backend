package app

import (
	"user-service/internal/database"

	"user-service/internal/transport/handler"

	"log"
	"net/http"

	"github.com/gorilla/mux"
)



func Run() {
    database.Connect()
	r := mux.NewRouter()
	r.HandleFunc("/", handler.HelloWorld).Methods("GET")
	r.HandleFunc("/users", handler.CreateUserHandler).Methods("POST")
	r.HandleFunc("/users/{id:[0-9]+}", handler.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", handler.UpdateUserHandler).Methods("PUT")
	r.HandleFunc("/users/{id:[0-9]+}", handler.DeleteUserHandler).Methods("DELETE")

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}