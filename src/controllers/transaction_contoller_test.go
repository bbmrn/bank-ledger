package controllers

import (
	"bank-ledger/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// TestProcessTransaction tests the ProcessTransaction handler function.
// It verifies the following scenarios:
//  1. Valid credit transaction - Ensures a credit transaction is processed successfully
//  2. Valid debit transaction - Ensures a debit transaction is processed successfully
//  3. Insufficient balance - Verifies proper error handling when account has insufficient funds
//
// For each test case it:
//   - Sets up a mock HTTP request with transaction data
//   - Calls ProcessTransaction handler
//   - Validates the HTTP status code
//   - Validates the response body matches expected output
//
// The test uses table-driven testing pattern with Gin's test mode.
func TestProcessTransaction(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		transaction    models.Transaction
		expectedStatus int
		expectedBody   map[string]string
	}{
		{
			name: "Valid credit transaction",
			transaction: models.Transaction{
				UserID:      1,
				Amount:      100.0,
				Type:        "credit",
				Description: "Test credit",
			},
			expectedStatus: http.StatusCreated,
			expectedBody: map[string]string{
				"message": "Transaction processed successfully",
			},
		},
		{
			name: "Valid debit transaction",
			transaction: models.Transaction{
				UserID:      1,
				Amount:      50.0,
				Type:        "debit",
				Description: "Test debit",
			},
			expectedStatus: http.StatusCreated,
			expectedBody: map[string]string{
				"message": "Transaction processed successfully",
			},
		},
		{
			name: "Insufficient balance",
			transaction: models.Transaction{
				UserID:      1,
				Amount:      1000000.0,
				Type:        "debit",
				Description: "Test insufficient",
			},
			expectedStatus: http.StatusBadRequest,
			expectedBody: map[string]string{
				"error": "Insufficient balance",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			// Create request body
			jsonData, _ := json.Marshal(tt.transaction)
			c.Request = httptest.NewRequest("POST", "/transaction", bytes.NewBuffer(jsonData))
			c.Request.Header.Set("Content-Type", "application/json")

			ProcessTransaction(c)

			assert.Equal(t, tt.expectedStatus, w.Code)

			var response map[string]string
			err := json.Unmarshal(w.Body.Bytes(), &response)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedBody, response)
		})
	}
}
