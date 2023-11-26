package services

import (
	"github.com/go-ddd-bank/domain/dto"
	domain "github.com/go-ddd-bank/domain/model"
	repo "github.com/go-ddd-bank/domain/repository"
	errors "github.com/go-ddd-bank/utils"
)

type UserAccountService struct {
	r repo.UserAccountRepository
}

func NewUserAccountService(r repo.UserAccountRepository) *UserAccountService {
	return &UserAccountService{r: r}
}

func (uas *UserAccountService) GetUserAccountDetails(userID int64) (*dto.UserAccountDto, *errors.Errors) {
	userAccountDto, err := uas.r.GetAccountViewByUserID(userID)
	if err != nil {
		return nil, err
	}

	return toAccountUserWithoutCreationDate(userAccountDto), nil
}

func toAccountUserWithoutCreationDate(userAcc *dto.UserAccountDto) *dto.UserAccountDto {
	return &dto.UserAccountDto{User: userAcc.User,
		Account: &domain.Account{
			ID:             userAcc.Account.ID,
			UserID:         userAcc.Account.UserID,
			AccountsNumber: userAcc.Account.AccountsNumber, Expenses: userAcc.Account.Expenses,
			Income:         userAcc.Account.Income,
			Balance:        userAcc.Account.Balance,
			CardNumber:     userAcc.Account.CardNumber,
			ExpirationDate: userAcc.Account.ExpirationDate,
		}}
}
