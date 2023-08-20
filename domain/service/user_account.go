package services

import (
	"github.com/go-ddd-bank/domain/dto"
	repo "github.com/go-ddd-bank/domain/repository"
	errors "github.com/go-ddd-bank/utils"
)

type UserAccountService struct {
	r repo.UserAccountRepository
}

func NewUserAccountService(r repo.UserAccountRepository) *UserAccountService {
	return &UserAccountService{r: r}
}

func (uar *UserAccountService) GetUserAccountDetails(userID int64) (*dto.UserAccountDto, *errors.Errors) {
	userAccountDto, err := uar.r.GetAccountViewByUserID(userID)
	if err != nil {
		return nil, err
	}

	return userAccountDto, nil
}
