package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbName         = "urlshortener"
	collectionName = "urls"
	expiryDuration = 24 * time.Hour // Adjust as per your requirements
)

func init() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	// Ping MongoDB to verify the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}

	collection = client.Database(dbName).Collection(collectionName)
}
