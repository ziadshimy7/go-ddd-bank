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
	queryGetAccountByID = `SELECT * FROM accounts WHERE user_id = ?;`
)

func NewAccountRepository(db *db.DbConnection) *AccountRepo {
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
	acc := &domain.Account{}

	err = result.Scan(&acc.ID, &acc.UserID,
		&acc.AccountsNumber, &acc.Expenses, &acc.Income,
		&acc.Balance, &acc.CardNumber, &acc.ExpirationDate, &acc.CreatedAt,
	)

	if err != nil {
		fmt.Println(err)
		return nil, errors.NewInternalServerError("Cannot fetch account details")
	}

	return acc, nil
}
