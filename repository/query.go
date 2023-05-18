package repository

import (
	"context"
	"errors"
	"time"

	"github.com/prashant3286/short-url-service/models"
	"github.com/prashant3286/short-url-service/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	collection *mongo.Collection
)

func ShortenURL(url *models.URL) (string, error) {
	// Generate a unique short URL
	shortURL := utils.GenerateRandomString(8)

	// Create a URL document
	urlDocument := bson.M{
		"_id":        shortURL,
		"long_url":   url.LongURL,
		"created_at": time.Now(),
		"expiry_at":  time.Now().Add(expiryDuration),
	}

	// Insert the URL document into the collection
	_, err := collection.InsertOne(context.Background(), urlDocument)
	if err != nil {
		return "", err
	}

	return shortURL, nil
}

func GetLongURL(shortURL string) (string, error) {
	var result bson.M

	// Find the URL document by short URL
	err := collection.FindOne(context.Background(), bson.M{"_id": shortURL}).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return "", errors.New("URL not found")
		}
		return "", err
	}

	// Check if the URL has expired
	expiryAt := result["expiry_at"].(primitive.DateTime).Time()
	if time.Now().After(expiryAt) {
		return "", errors.New("URL has expired")
	}

	return result["long_url"].(string), nil
}
