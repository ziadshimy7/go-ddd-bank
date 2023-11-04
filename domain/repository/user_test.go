package repo

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	domain "github.com/go-ddd-bank/domain/model"
	da "github.com/go-ddd-bank/infrastructure/db"
	"github.com/stretchr/testify/assert"
)

func TestGetByEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	r := NewUserRepository(&da.DbConnection{DB: db})

	mockUser := &domain.User{Email: "Ziadshimy7@gmail.com"}

	rows := sqlmock.NewRows([]string{
		"ID", "FirstName", "LastName", "Password", "Phone",
	}).
		AddRow(1, "Ziad", "Elshimy", "Cocowawa_12345!", "+79122390087")

	mock.ExpectPrepare(regexp.QuoteMeta(queryGetUserByEmail)).
		ExpectQuery().
		WithArgs("Ziadshimy7@gmail.com").
		WillReturnRows(rows)

	customErr := r.GetByEmail(mockUser)

	assert.Nil(t, customErr)
	assert.NotNil(t, mockUser.FirstName)
	assert.NotNil(t, mockUser.ID)
	assert.NotNil(t, mockUser.Password)
	assert.NotNil(t, mockUser.Phone)
	assert.NotNil(t, mockUser.LastName)
	assert.Equal(t, mockUser.FirstName, "Ziad")
	assert.NoError(t, mock.ExpectationsWereMet())

}
