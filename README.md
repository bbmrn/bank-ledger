# API Documentation

## Base URL
`http://localhost:8080`

## 1. Account Management

### Create a New Account
Create a new user account.

**Endpoint:** `POST /api/accounts`

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

**Endpoint:** `GET /api/accounts/:id`

#### Responses
**Success (200 OK)**
{
   "id": 1,
   "name": "John Doe",
   "email": "john@example.com",
   "balance": 1000.0,
   "created_at": "2023-09-07T12:34:56Z"
}


**Error Responses**
- 400 Bad Request: `{"error": "Invalid ID"}`
- 404 Not Found: `{"error": "Account not found"}`
- 500 Internal Server Error: `{"error": "Failed to fetch account details"}`

## 2. Transaction Management

### Process a Transaction
Process a transaction (credit or debit) for a user.

**Endpoint:** `POST /api/transactions`

#### Request Body
{
   "account_id": 1,
   "amount": 100.0,
   "type": "credit",
   "description": "Deposit"
}

#### Responses
**Success (201 Created)**
{
   "id": "tx123",
   "account_id": 1,
   "amount": 100.0,
   "type": "credit",
   "description": "Deposit",
   "created_at": "2023-09-07T12:34:56Z"
}

**Error Responses**
- 400 Bad Request: `{"error": "Invalid request payload"}`
- 404 Not Found: `{"error": "Account not found"}`
- 500 Internal Server Error: `{"error": "Failed to process transaction"}`

### Get Transaction History
Retrieve the transaction history for a specific account.

**Endpoint:** `GET /api/transactions/:account_id`

#### Responses
**Success (200 OK)**
{
   "account_id": 1,
   "transactions": [
      {
         "id": "tx123",
         "account_id": 1,
         "amount": 100.0,
         "type": "credit",
         "description": "Deposit",
         "created_at": "2023-09-07T12:34:56Z"
      }
   ]
}


**Error Responses**
- 400 Bad Request: `{"error": "Invalid account ID"}`
- 404 Not Found: `{"error": "Account not found"}`
- 500 Internal Server Error: `{"error": "Failed to fetch transaction history"}`

## 3. Example Requests

### Create Account
curl -X POST http://localhost:8080/api/accounts \
   -H "Content-Type: application/json" \
   -d '{
      "name": "John Doe",
      "email": "john@example.com",
      "balance": 1000.0
   }'


### Process Transaction
curl -X POST http://localhost:8080/api/transactions \
   -H "Content-Type: application/json" \
   -d '{
      "account_id": 1,
      "amount": 100.0,
      "type": "credit",
      "description": "Deposit"
   }'

## 4. Development

### Prerequisites
- Go 1.19 or higher
- PostgreSQL
- Make

### Running Tests
make test

### Running the Application
make run

