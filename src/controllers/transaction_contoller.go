package controllers

import (
	"net/http"

	"bank-ledger/database"
	"bank-ledger/models"

	"github.com/gin-gonic/gin"
)

// ProcessTransaction handles financial transactions for user accounts.
// It processes both debit and credit transactions while ensuring atomic operations
// using database transactions.
//
// The function performs the following steps:
// 1. Validates the incoming JSON transaction request
// 2. Begins a database transaction
// 3. Locks and retrieves the current user balance
// 4. Validates sufficient balance for debit transactions
// 5. Updates the user's balance
// 6. Logs the transaction details
// 7. Commits the database transaction
//
// Parameters:
//   - c *gin.Context: The Gin context containing the HTTP request and response
//
// Request Body:
//   - UserID: The ID of the user performing the transaction
//   - Amount: The transaction amount
//   - Type: The transaction type ("debit" or "credit")
//   - Description: A description of the transaction
//
// Returns:
//   - 201 StatusCreated: Transaction processed successfully
//   - 400 StatusBadRequest: Invalid request payload or insufficient balance
//   - 500 StatusInternalServerError: Database or transaction processing errors
func ProcessTransaction(c *gin.Context) {
	var transaction models.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Start a database transaction
	tx, err := database.BeginTransaction()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start transaction"})
		return
	}
	defer tx.Rollback()

	// Update account balance
	var balance float64
	err = tx.QueryRow("SELECT balance FROM users WHERE id = $1 FOR UPDATE", transaction.UserID).Scan(&balance)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch balance"})
		return
	}

	if transaction.Type == "debit" && balance < transaction.Amount {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Insufficient balance"})
		return
	}

	if transaction.Type == "debit" {
		balance -= transaction.Amount
	} else {
		balance += transaction.Amount
	}

	_, err = tx.Exec("UPDATE users SET balance = $1 WHERE id = $2", balance, transaction.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update balance"})
		return
	}

	// Insert transaction log
	_, err = tx.Exec("INSERT INTO transactions (user_id, amount, type, description) VALUES ($1, $2, $3, $4)", transaction.UserID, transaction.Amount, transaction.Type, transaction.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to log transaction"})
		return
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to commit transaction"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Transaction processed successfully"})
}
