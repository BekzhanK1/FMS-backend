package app

import (
	"context"
	db "document-service/internal/database"
	store "document-service/internal/repository"
	"document-service/internal/service"
	httpHandler "document-service/internal/transport/http"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Run() {
	mongoClient, err := db.MongoConnect()
	if err != nil {
		log.Fatalf("Error occured connecting to Mongo: %s", err)
	}
	defer mongoClient.Disconnect(context.TODO())

	mongoDb := mongoClient.Database("fms")
	
	documentStore := store.NewStore(mongoDb)
	documentService := service.NewService(documentStore)
	documentHandler := httpHandler.NewHanlder(*documentService)

	r := mux.NewRouter()

	documentRouter := r.PathPrefix("/documents").Subrouter()
	documentHandler.RegisterRoutes(documentRouter)

	log.Println("Starting server on :8081")
	err = http.ListenAndServe(":8081", r)
	if err != nil {
		log.Fatal(err)
	}
}