package domain

import (
	"strings"
	"time"

	errors "github.com/go-ddd-bank/utils"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int64      `json:"ID,,omitempty"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Password  string     `json:"password,omitempty"`
	Email     string     `json:"email"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
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
	}
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

func (u *User) VerifyPassword(userPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(userPassword))
	return err != nil
}
