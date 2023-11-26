package domain

import "time"

type Account struct {
	ID             int64      `json:"accountId,omitempty"`
	UserID         int64      `json:"userId,omitempty"`
	AccountsNumber string     `json:"accountNumber"`
	Expenses       float32    `json:"expenses"`
	Income         float32    `json:"income"`
	Balance        float32    `json:"balance"`
	CardNumber     string     `json:"cardNumber,omitempty"`
	ExpirationDate *time.Time `json:"expirationDate"`
	CreatedAt      *time.Time `json:"createdAt,omitempty"`
}
