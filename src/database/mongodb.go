// database/package.go
package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client         *mongo.Client
	defaultTimeout = 10 * time.Second // Default timeout for database operations
)

func Init() {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err = mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
}

func Disconnect() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := client.Disconnect(ctx); err != nil {
		log.Fatalf("Failed to disconnect MongoDB client: %v", err)
	}
}

// Example of a function using the default timeout
func GetMongoCollection(collectionName string) *mongo.Collection {
	return client.Database(os.Getenv("DATABASE_NAME")).Collection(collectionName)
}

func Find[T any](collectionName string, condition bson.D) []T {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	collection := GetMongoCollection(collectionName)
	cursor, err := collection.Find(ctx, condition)
	if err != nil {
		panic(err)
	}
	defer cursor.Close(ctx)

	var items []T
	if err = cursor.All(ctx, &items); err != nil {
		panic(err)
	}

	return items
}
