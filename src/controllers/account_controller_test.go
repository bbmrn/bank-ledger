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
