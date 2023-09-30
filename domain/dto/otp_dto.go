package dto

type OTPDto struct {
	ID           string `json:"id"`
	FirstName    string `json:"name"`
	Email        string `json:"email"`
	Otp_Enabled  bool   `json:"otp_enabled"`
	Otp_verified bool   `json:"otp_verified"`
}
