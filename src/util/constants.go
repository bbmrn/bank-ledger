package util

const (
	// Database and Collection Names
	MongoDBName           = "innoscripta"
	TransactionCollection = "transactions"

	// RabbitMQ Configuration
	RabbitMQQueueName = "transactions"
	RabbitMQURL       = "amqp://guest:guest@localhost:5672/"

	// PostgreSQL Configuration
	PostgreSQLURL = "user=youruser dbname=yourdb sslmode=disable password=yourpassword"

	// API Endpoints
	APIVersion           = "/api/v1"
	AccountsEndpoint     = APIVersion + "/accounts"
	TransactionsEndpoint = APIVersion + "/transactions"

	// Error Messages
	ErrUserNotFound   = "user not found"
	ErrInvalidID      = "invalid ID"
	ErrInternalServer = "internal server error"
)
