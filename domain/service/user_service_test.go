package services

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	da "github.com/go-ddd-bank/infrastructure/db"

	"github.com/go-ddd-bank/domain/dto"
	domain "github.com/go-ddd-bank/domain/model"
	repo "github.com/go-ddd-bank/domain/repository"
	test_utils "github.com/go-ddd-bank/utils/tests"
	"github.com/stretchr/testify/assert"
)

const queryCreateUser = "INSERT INTO users(first_name,last_name,email,password,phone) VALUES (?,?,?,?,?);"

func TestUserServiceCreateUser(t *testing.T) {
	db, mock, err := test_utils.InitTestingDatabaseAdapter()
	assert.NoError(t, err)
	defer db.Close()

	r := repo.NewUserRepository(&da.DbConnection{DB: db})

	mock.ExpectPrepare(regexp.QuoteMeta(queryCreateUser)).
		ExpectExec().
		WithArgs("Ziad", "Elshimy", "Ziadshimy7@gmail.com", sqlmock.AnyArg(), "+79122390087").WillReturnResult(sqlmock.NewResult(1, 1))

	us := NewUserService(r)

	user := domain.NewUser(&domain.User{Email: "Ziadshimy7@gmail.com", ID: 1,
		FirstName: "Ziad", LastName: "Elshimy", Password: "asdasdasd", Phone: "+79122390087"})

	mockUser, customErr := us.CreateUser(user)

	t.Run("Should successfully create a user and return it without password", func(t *testing.T) {
		expectedType := &dto.UserDTO{}

		assert.NotNil(t, mockUser)
		assert.Nil(t, customErr)
		assert.NotNil(t, mockUser.Email)
		assert.NotNil(t, mockUser.FirstName)
		assert.NotNil(t, mockUser.LastName)
		assert.NotNil(t, mockUser.ID)
		assert.NotNil(t, mockUser.Phone)
		assert.IsType(t, expectedType, mockUser)
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}
