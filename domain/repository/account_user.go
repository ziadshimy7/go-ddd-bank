package repo

import (
	"database/sql"

	"github.com/go-ddd-bank/domain/dto"
	domain "github.com/go-ddd-bank/domain/model"
	"github.com/go-ddd-bank/infrastructure/db"
	errors "github.com/go-ddd-bank/utils"
)

const queryGetUserAccountByID = `
SELECT accounts_id, user_id, account_number, expenses, income, balance,
       first_name, last_name, email FROM accounts 
       INNER JOIN users ON users.id = accounts.user_id WHERE accounts.user_id = ?;
`

type UserAccountRepository interface {
	GetAccountViewByUserID(userID int64) (*dto.UserAccountDto, *errors.Errors)
}

type UserAccountDtoRepository struct {
	Db *sql.DB
}

func NewAccountUserViewRepository(db *db.Adapter) *UserAccountDtoRepository {
	return &UserAccountDtoRepository{Db: db.DB}
}

func (r *UserAccountDtoRepository) GetAccountViewByUserID(userID int64) (*dto.UserAccountDto, *errors.Errors) {
	dto := &dto.UserAccountDto{
		Account: &domain.Account{},
		User:    &domain.User{},
	}

	stmt, err := r.Db.Prepare(queryGetUserAccountByID)

	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}

	result := stmt.QueryRow(userID)

	err = result.Scan(&dto.Account.ID, &dto.Account.UserID,
		&dto.Account.AccountsNumber, &dto.Account.Expenses, &dto.Account.Income,
		&dto.Account.Balance, &dto.User.FirstName,
		&dto.User.LastName, &dto.User.Email)

	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}

	return dto, nil
}
