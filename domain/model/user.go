package domain

import (
	"strings"
	"time"

	errors "github.com/go-ddd-bank/utils"
)

type User struct {
	ID        int64      `json:"ID,,omitempty"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Password  string     `json:"password,omitempty"`
	Email     string     `json:"email"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`

	Otp_enabled  bool `json:"otp_enabled"`
	Otp_verified bool `json:"otp_verified"`

	Otp_secret   string `json:"otp_secret,omitempty"`
	Otp_auth_url string `json:"otp_auth_url,omitempty"`
}

type OTPInput struct {
	UserId string `json:"user_id"`
	Token  string `json:"token"`
}

func (user *User) Validate() *errors.Errors {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Email = strings.TrimSpace(user.Email)
	if user.Email == "" {
		errors.NewBadRequestError("Email cannot be empty")
	}
	user.Password = strings.TrimSpace(user.Password)

	if user.Password == "" {
		errors.NewBadRequestError("Invalid password")
	}

	return nil
}
