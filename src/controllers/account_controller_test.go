package controllers

import (
	"bank-ledger/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// TestCreateAccount tests the CreateAccount handler function
// It verifies that:
// - A new account can be created via POST request to "/accounts"
// - The response status code is 201 (Created)
// - The response body contains the created user with valid ID
// - The created user matches the test user data (name, email, balance)
// The test:
// 1. Sets up a test Gin router
// 2. Creates test user data
// 3. Makes a POST request with the test user JSON
// 4. Validates the response status and user data
func TestCreateAccount(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/accounts", CreateAccount)

	testUser := models.User{
		Name:      "Test User",
		Email:     "test@example.com",
		Balance:   100.00,
		CreatedAt: time.Now(),
	}

	jsonValue, _ := json.Marshal(testUser)
	req, _ := http.NewRequest("POST", "/accounts", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response models.User
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.NotEmpty(t, response.ID)
	assert.Equal(t, testUser.Name, response.Name)
	assert.Equal(t, testUser.Email, response.Email)
	assert.Equal(t, testUser.Balance, response.Balance)
}

func TestGetAccount(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/accounts/:id", GetAccount)

	// Test existing user
	req, _ := http.NewRequest("GET", "/accounts/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	// Test non-existent user
	req, _ = http.NewRequest("GET", "/accounts/999", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}
