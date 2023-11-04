package repo

import (
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	domain "github.com/go-ddd-bank/domain/model"
	da "github.com/go-ddd-bank/infrastructure/db"
	"github.com/stretchr/testify/assert"
)

func TestGetAccount(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	r := NewAccountRepository(&da.DbConnection{DB: db})

	mockAccount := &domain.Account{ID: 1}

	rows := sqlmock.NewRows([]string{
		"ID", "UserID", "AccountsNumber", "Expenses", "Income",
		"Balance", "CardNumber", "ExpirationDate", "CreatedAt",
	}).
		AddRow(1, 1, "123456", 100.5, 200.5, 100, "1234-5678-1234-5678", time.Now(), time.Now())

	mock.ExpectPrepare(regexp.QuoteMeta(queryGetAccountByID)).
		ExpectQuery().
		WithArgs(1).
		WillReturnRows(rows)

	account, customErr := r.GetAccount(mockAccount)

	assert.NoError(t, err)
	assert.Nil(t, customErr)
	assert.Equal(t, mockAccount.ID, account.ID)

	assert.NoError(t, mock.ExpectationsWereMet())
}
