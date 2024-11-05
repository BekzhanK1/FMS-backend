package app

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	db "document-service/internal/database"
	store "document-service/internal/repository"
	"document-service/internal/service"
	handler "document-service/internal/transport/rpc"
)

func Run() {
	mongoClient, err := db.MongoConnect()
	if err != nil {
		log.Fatalf("Error occured connecting to Mongo: %s", err)
	}
	defer mongoClient.Disconnect(context.TODO())

	mongoDb := mongoClient.Database("fms")

	lis, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatalf("ERROR STARTING THE SERVER : %v", err)
	}

	grpcServer := grpc.NewServer()

	documentStore := store.NewStore(mongoDb)
	documentService := service.NewService(documentStore)

	handler.NewServer(grpcServer, documentService)

	log.Println("Starting gprc server on :5001")
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
