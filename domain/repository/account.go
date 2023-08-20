package repo

import (
	"database/sql"
	"fmt"
	domain "github.com/go-ddd-bank/domain/model"
	"github.com/go-ddd-bank/infrastructure/db"
	errors "github.com/go-ddd-bank/utils"
)

type AccountRepository interface {
	GetAccount(*domain.Account) (*domain.Account, *errors.Errors)
}

type AccountRepo struct {
	Db *sql.DB
}

var (
	queryGetAccountByID = `SELECT accounts_id, user_id, account_number, expenses, income, balance
	 FROM accounts WHERE accounts_id = ?;`
)

func NewAccountRepository(db *db.Adapter) *AccountRepo {
	return &AccountRepo{Db: db.DB}
}

func (r *AccountRepo) GetAccount(account *domain.Account) (*domain.Account, *errors.Errors) {
	stmt, err := r.Db.Prepare(queryGetAccountByID)

	if err != nil {
		fmt.Println(err)
		return nil, errors.NewInternalServerError("Cannot fetch account details")
	}

	defer stmt.Close()

	result := stmt.QueryRow(account.ID)
	userAcc := &domain.Account{}

	err = result.Scan(&userAcc.ID, &userAcc.UserID,
		&userAcc.AccountsNumber, &userAcc.Expenses, &userAcc.Income,
		&userAcc.Balance,
	)

	if err != nil {
		fmt.Println(err)
		return nil, errors.NewInternalServerError("Cannot fetch account details")
	}

	return userAcc, nil
}
