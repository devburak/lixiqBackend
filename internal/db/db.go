// mongo.go

package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var database *mongo.Database

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}

	// .env get MONGO URI
	mongoURI := os.Getenv("MONGO_URI")
	mongoDBName := os.Getenv("MONGO_DB_NAME")
	if mongoURI == "" {
		log.Fatal("MONGO_URI is not set in .env file")
	}

	// MongoDB connection created
	clientOptions := options.Client().ApplyURI(mongoURI) // MongoDB URL'nizi burada belirtin

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %s", err)
	}

	// MongoDB ping status
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %s", err)
	}

	database = client.Database(mongoDBName) // MongoDB veritabanı adınızı buraya yazın

	fmt.Printf("connection is OK")
}

func Close() {
	if client != nil {
		client.Disconnect(context.Background())
	}
}

// Database variable as a function
func GetDatabase() *mongo.Database {
	return database
}
