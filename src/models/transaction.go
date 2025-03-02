package models

import "time"

type Transaction struct {
	ID          string    `json:"id" bson:"_id"`
	UserID      int       `json:"user_id" bson:"user_id"`
	Amount      float64   `json:"amount" bson:"amount"`
	Type        string    `json:"type" bson:"type"` // "credit" or "debit"
	Description string    `json:"description" bson:"description"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
}
