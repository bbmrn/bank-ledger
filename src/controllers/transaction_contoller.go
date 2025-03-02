package controllers

import (
	"net/http"

	"bank-ledger/database"
	"bank-ledger/models"

	"github.com/gin-gonic/gin"
)

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
