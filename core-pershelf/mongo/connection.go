package mongo

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/core-pershelf/globals"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongoDB() error {
	fmt.Println("🔧 Starting MongoDB connection...")

	// Determine the path to the .env file
	envPath := filepath.Join("..", "..", ".env")

	// Load the .env file
	if err := godotenv.Load(envPath); err != nil {
		fmt.Printf("❌ Error loading .env file from %s: %v\n", envPath, err)
		return fmt.Errorf("error loading .env file: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Get the mongo uri from the .env file
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		return fmt.Errorf("MONGO_URI environment variable is not set")
	}
	fmt.Printf("🔗 Mongo URI: %s\n", mongoURI)
	clientOptions := options.Client().ApplyURI(mongoURI)

	fmt.Println("📡 Attempting to connect...")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Printf("❌ MongoDB connection failed: %v\n", err)
		return fmt.Errorf("MongoDB connection failed: %v", err)
	}

	fmt.Println("📶 Attempting ping to MongoDB server...")
	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Printf("❌ MongoDB ping failed: %v\n", err)
		return fmt.Errorf("MongoDB ping failed: %v", err)
	}

	globals.MongoClient = client
	fmt.Println("✅ MongoDB connection successful!")
	return nil
}
