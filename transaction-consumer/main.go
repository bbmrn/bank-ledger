package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"bank-ledger/database"
	"bank-ledger/models"

	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

func initMongoDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	var err error
	mongoClient, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	err = mongoClient.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Error pinging MongoDB: %v", err)
	}

	fmt.Println("Successfully connected to MongoDB!")
}

func main() {
	// Initialize databases
	database.InitPostgres()
	initMongoDB()

	// Connect to RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Error connecting to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Error opening RabbitMQ channel: %v", err)
	}
	defer ch.Close()

	// Declare the queue
	q, err := ch.QueueDeclare(
		"transactions", // Queue name
		false,          // Durable
		false,          // Delete when unused
		false,          // Exclusive
		false,          // No-wait
		nil,            // Arguments
	)
	if err != nil {
		log.Fatalf("Error declaring queue: %v", err)
	}

	// Consume messages
	msgs, err := ch.Consume(
		q.Name, // Queue
		"",     // Consumer
		true,   // Auto-ack
		false,  // Exclusive
		false,  // No-local
		false,  // No-wait
		nil,    // Args
	)
	if err != nil {
		log.Fatalf("Error registering consumer: %v", err)
	}

	log.Println("Waiting for transactions...")

	// Process messages
	for msg := range msgs {
		var transaction models.Transaction
		err := json.Unmarshal(msg.Body, &transaction)
		if err != nil {
			log.Printf("Error decoding transaction: %v", err)
			continue
		}

		// Insert transaction into MongoDB
		collection := mongoClient.Database("innoscripta").Collection("transactions")
		_, err = collection.InsertOne(context.Background(), transaction)
		if err != nil {
			log.Printf("Error inserting transaction into MongoDB: %v", err)
		}

		log.Printf("Processed transaction: %+v", transaction)
	}
}
