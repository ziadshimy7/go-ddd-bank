package domain

type Account struct {
	ID             int64   `json:"accounts_id,omitempty"`
	UserID         int64   `json:"user_id,omitempty"`
	AccountsNumber string  `json:"accounts_number"`
	Expenses       float32 `json:"expenses"`
	Income         float32 `json:"income"`
	Balance        float32 `json:"balance"`
}