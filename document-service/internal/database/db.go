package db

import (
	"context"
	"document-service/internal/config"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnect() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(config.Envs.MongoUri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	if err = client.Ping(context.TODO(), nil); err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	fmt.Println("Connected to mongo successfully!")
	return client, nil
}
