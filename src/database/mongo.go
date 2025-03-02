package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"bank-ledger/util"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var TransactionCollection *mongo.Collection

// InitMongoDB initializes a connection to MongoDB using environment configurations.
// It performs the following steps:
// 1. Creates a context with 10 second timeout
// 2. Establishes connection to MongoDB using MONGO_URL from environment (defaults to localhost:27017)
// 3. Verifies connection by pinging the database
// 4. Initializes the transaction collection reference
//
// The function will terminate the program with a fatal error if:
// - Unable to establish connection to MongoDB
// - Unable to ping the MongoDB server
//
// Global variables set:
// - MongoClient: The MongoDB client connection
// - TransactionCollection: Reference to the transaction collection
func InitMongoDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(util.GetEnv("MONGO_URL", "mongodb://localhost:27017"))
	var err error
	MongoClient, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	err = MongoClient.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Error pinging MongoDB: %v", err)
	}

	TransactionCollection = MongoClient.Database(util.MongoDBName).Collection(util.TransactionCollection)
	fmt.Println("Successfully connected to MongoDB!")
}
