package domain

import "time"

type OTP struct {
	ID           int64      `json:"otpId"`
	Otp_enabled  bool       `json:"otpEnabled"`
	Otp_verified bool       `json:"otpVerified"`
	Otp_secret   string     `json:"otpSecret,omitempty"`
	Otp_auth_url string     `json:"otpAuthUrl,omitempty"`
	Passcode     string     `json:"otpPasscode,omitempty"`
	User_id      int64      `json:"userId"`
	CreatedAt    *time.Time `json:"createdAt,omitempty"`
	UpdatedAt    *time.Time `json:"updatedAt,omitempty"`
}

type OTPInput struct {
	UserId string `json:"userId"`
	Token  string `json:"token"`
}

type OTPTokenRequest struct {
	Token string `json:"token"`
}

func (otp *OTP) Validate(token string) bool {
	return otp.Passcode == token
}
