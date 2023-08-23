package domain

import "time"

type OTP struct {
	ID           int64      `json:"otp_id"`
	Otp_enabled  bool       `json:"otp_enabled"`
	Otp_verified bool       `json:"otp_verified"`
	Otp_secret   string     `json:"otp_secret,omitempty"`
	Otp_auth_url string     `json:"otp_auth_url,omitempty"`
	User_id      int64      `json:"user_id"`
	CreatedAt    *time.Time `json:"created_at,omitempty"`
	UpdatedAt    *time.Time `json:"updated_at,omitempty"`
}

type OTPInput struct {
	UserId string `json:"user_id"`
	Token  string `json:"token"`
}

type OTPTokenRequest struct {
	Token string `json:"token"`
}
