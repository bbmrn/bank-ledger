package controllers

import (
	"database/sql"
	"net/http"

	"bank-ledger/database"
	"bank-ledger/models"

	"github.com/gin-gonic/gin"
)

// CreateAccount handles account creation
// CreateAccount godoc
// @Summary Create a new bank account
// @Description Creates a new user account with the provided details
// @Tags accounts
// @Accept json
// @Produce json
// @Param user body models.User true "User account information"
// @Success 201 {object} models.User "Successfully created account"
// @Failure 400 {object} gin.H "Bad request - Invalid input"
// @Failure 500 {object} gin.H "Internal server error"
// @Router /accounts [post]
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
// GetAccount retrieves a single user account from the database by ID.
// It responds with the user details including ID, name, email, balance, and creation timestamp.
//
// @Summary Get user account details
// @Description Fetch a user account by ID from the database
// @Tags accounts
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} models.User
// @Failure 404 {object} object "User not found"
// @Failure 500 {object} object "Internal server error"
// @Router /accounts/{id} [get]
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
