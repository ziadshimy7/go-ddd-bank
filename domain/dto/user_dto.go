package dto

type UserDTO struct {
	ID        int64  `json:"id,omitempty"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}
