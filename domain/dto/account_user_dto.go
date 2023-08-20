package dto

import (
	domain "github.com/go-ddd-bank/domain/model"
)

type UserAccountDto struct {
	Account *domain.Account `json:"account"`
	User    *domain.User    `json:"user"`
}
