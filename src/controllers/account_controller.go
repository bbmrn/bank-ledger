package controllers

import (
	"database/sql"
	"net/http"

	"bank-ledger/database"
	"bank-ledger/models"

	"github.com/gin-gonic/gin"
)

// CreateAccount handles account creation
func CreateAccount(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `INSERT INTO users (name, email, balance, created_at) VALUES ($1, $2, $3, $4) RETURNING id`
	err := database.PostgresDB.QueryRow(query, user.Name, user.Email, user.Balance, user.CreatedAt).Scan(&user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// GetAccount handles fetching account details
func GetAccount(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	query := `SELECT id, name, email, balance, created_at FROM users WHERE id = $1`
	err := database.PostgresDB.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.Balance, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, user)
}
