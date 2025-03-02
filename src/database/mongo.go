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
