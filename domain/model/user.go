package domain

import (
	"strings"
	"time"

	errors "github.com/go-ddd-bank/utils"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int64      `json:"id,omitempty"`
	FirstName string     `json:"firstName"`
	LastName  string     `json:"lastName"`
	Password  string     `json:"password,omitempty"`
	Email     string     `json:"email"`
	Phone     string     `json:"phone"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

func NewUser(args *User) *User {
	return &User{
		ID:        args.ID,
		FirstName: args.FirstName,
		LastName:  args.LastName,
		Password:  args.Password,
		Email:     args.Email,
		CreatedAt: args.CreatedAt,
		UpdatedAt: args.UpdatedAt,
		Phone:     args.Phone,
	}
}

func (user *User) Validate() *errors.Errors {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Email = strings.TrimSpace(user.Email)
	user.Phone = strings.TrimSpace(user.Phone)
	user.Password = strings.TrimSpace(user.Password)

	if user.Email == "" {
		return errors.NewBadRequestError("Email cannot be empty")
	}

	if user.Password == "" {
		return errors.NewBadRequestError("Password cannot be empty")
	}

	if user.Phone == "" {
		return errors.NewBadRequestError("Phone cannot be empty")
	}

	if user.FirstName == "" {
		return errors.NewBadRequestError("First name cannot be empty")
	}

	if user.LastName == "" {
		return errors.NewBadRequestError("Last name cannot be empty")
	}

	return nil
}

func (u *User) VerifyPassword(userPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(userPassword))
	return err != nil
}

func (u *User) HashPassword() (string, *errors.Errors) {
	pwSlice, err := bcrypt.GenerateFromPassword([]byte(u.Password), 12)
	if err != nil {
		return "", errors.NewBadRequestError("Failed to encrypt the pw")
	}
	return string(pwSlice), nil
}
