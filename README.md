# API Documentation

## Base URL
`http://localhost:8080`

## 1. Account Management

### Create a New Account
Create a new user account.

**Endpoint:** `POST /accounts`

#### Request Body
{
   "name": "John Doe",
   "email": "john@example.com",
   "balance": 1000.0
}

#### Responses
**Success (201 Created)**
{
   "id": 1,
   "name": "John Doe",
   "email": "john@example.com",
   "balance": 1000.0,
   "created_at": "2023-09-07T12:34:56Z"
}

**Error Responses**
- 400 Bad Request: `{"error": "Invalid request payload"}`
- 500 Internal Server Error: `{"error": "Failed to create account"}`

### Get Account Details
Retrieve details of a specific account by ID.

**Endpoint:** `GET /accounts/:id`

#### Responses
**Success (200 OK)**
{
   "id": 1,
   "name": "John Doe",
   "email": "john@example.com",
   "balance": 1000.0,
   "created_at": "2023-09-07T12:34:56Z"
}
```

**Error Responses**
- 400 Bad Request: `{"error": "Invalid ID"}`
- 404 Not Found: `{"error": "User not found"}`
- 500 Internal Server Error: `{"error": "Failed to fetch account details"}`

## 2. Transaction Management

### Process a Transaction
Process a transaction (credit or debit) for a user.

**Endpoint:** `POST /transactions`

#### Request Body
{
   "user_id": 1,
   "amount": 100.0,
   "type": "credit",
   "description": "Deposit"
}

#### Responses
**Success (201 Created)**
{
   "message": "Transaction processed successfully"
}

**Error Responses**
- 400 Bad Request: `{"error": "Invalid request payload"}`
- 404 Not Found: `{"error": "User not found"}`
- 500 Internal Server Error: `{"error": "Failed to process transaction"}`

### Get Transaction History
Retrieve the transaction history (ledger) for a specific user.

**Endpoint:** `GET /transactions/history/:id`

#### Responses
**Success (200 OK)**
{
   "user_id": 1,
   "transactions": [
      {
         "id": "64f8a1b2e4b0c1a2b3c4d5e6",
         "user_id": 1,
         "amount": 100.0,
         "type": "credit",
         "description": "Deposit",
         "created_at": "2023-09-07T12:34:56Z"
      }
   ]
}


**Error Responses**
- 400 Bad Request: `{"error": "Invalid ID"}`
- 500 Internal Server Error: `{"error": "Failed to fetch transaction history"}`

## 3. Environment Variables

| Variable Name | Description | Default Value |
|--------------|-------------|---------------|
| POSTGRES_URL | PostgreSQL connection string | `user=youruser dbname=yourdb sslmode=disable password=yourpassword` |
| MONGO_URL | MongoDB connection string | `mongodb://localhost:27017` |
| RABBITMQ_URL | RabbitMQ connection string | `amqp://guest:guest@localhost:5672/` |

## 4. Example Requests

### Create Account
curl -X POST http://localhost:8080/accounts \
   -H "Content-Type: application/json" \
   -d '{
      "name": "John Doe",
      "email": "john@example.com",
      "balance": 1000.0
   }'

### Process Transaction
curl -X POST http://localhost:8080/transactions \
   -H "Content-Type: application/json" \
   -d '{
      "user_id": 1,
      "amount": 100.0,
      "type": "credit",
      "description": "Deposit"
   }'

## 5. Testing
Run unit tests using:
go test ./...

