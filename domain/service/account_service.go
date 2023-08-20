package services

import (
	"github.com/go-ddd-bank/domain/model"
	repo "github.com/go-ddd-bank/domain/repository"
	errors "github.com/go-ddd-bank/utils"
)

type AccountService struct {
	r *repo.AccountRepo
}

func NewAccountService(r *repo.AccountRepo) *AccountService {
	return &AccountService{r: r}
}

func (as *AccountService) GetAccountByID(userId int64) (*domain.Account, *errors.Errors) {
	account := &domain.Account{ID: userId}

	result, err := as.r.GetAccount(account)
	if err != nil {
		return nil, err
	}

	return result, nil
}
