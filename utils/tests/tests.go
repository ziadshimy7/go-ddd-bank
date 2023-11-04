package test_utils

import (
	"database/sql"
	"database/sql/driver"
	"math/rand"
	"strings"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	repo "github.com/go-ddd-bank/domain/repository"
	da "github.com/go-ddd-bank/infrastructure/db"
)

type MockOptions struct {
	Fields      []string
	FieldValues []string
	Query       string
	Args        []string
}

func NewMockOptions(fields, fieldValues []string, query string, args ...string) *MockOptions {
	return &MockOptions{
		Fields:      fields,
		FieldValues: fieldValues,
		Query:       query,
		Args:        args,
	}
}

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.NewSource(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder

	k := len(alphabet)

	for i := 0; i < n; i++ {

		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func InitTestingDatabaseAdapter() (*sql.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	return db, mock, err

}

func getDriverValue(args ...driver.Value) []driver.Value {
	return args
}

func MockUserRepositoryWithOptions(options MockOptions) (*repo.UserRepo, *sqlmock.Sqlmock) {
	db, mock, _ := InitTestingDatabaseAdapter()

	r := repo.NewUserRepository(&da.DbConnection{DB: db})

	rows := sqlmock.NewRows(options.Fields).
		AddRow(options.FieldValues)

	mock.ExpectPrepare(options.Query).
		ExpectQuery().
		WithArgs(options.Args).
		WillReturnRows(rows)

	return r, &mock
}
