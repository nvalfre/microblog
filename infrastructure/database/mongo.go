package database

import (
	"context"
	"microblog/infrastructure/logger"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// NewMongoClient initializes a new MongoDB client.
func NewMongoClient(uri string) *mongo.Client {
	opts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		logger.Error("Failed to create MongoDB client", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := client.Ping(ctx, nil); err != nil {
		logger.Error("Failed to connect to MongoDB", err)
	}
	return client
}
