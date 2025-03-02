package queues

import (
	"fmt"
	"log"

	"bank-ledger/util"

	"github.com/streadway/amqp"
)

var RabbitMQConn *amqp.Connection
var RabbitMQChannel *amqp.Channel

func InitRabbitMQ() {
	var err error
	RabbitMQConn, err = amqp.Dial(util.RabbitMQURL)
	if err != nil {
		log.Fatalf("Error connecting to RabbitMQ: %v", err)
	}

	RabbitMQChannel, err = RabbitMQConn.Channel()
	if err != nil {
		log.Fatalf("Error opening RabbitMQ channel: %v", err)
	}

	_, err = RabbitMQChannel.QueueDeclare(
		util.RabbitMQQueueName, // Queue name
		false,                  // Durable
		false,                  // Delete when unused
		false,                  // Exclusive
		false,                  // No-wait
		nil,                    // Arguments
	)
	if err != nil {
		log.Fatalf("Error declaring RabbitMQ queue: %v", err)
	}

	fmt.Println("Successfully connected to RabbitMQ!")
}
