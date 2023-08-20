package repo

import (
	"database/sql"
	"fmt"

	domain "github.com/go-ddd-bank/domain/model"
	"github.com/go-ddd-bank/infrastructure/db"
	errors "github.com/go-ddd-bank/utils"
)

type UserRepository interface {
	Save(user *domain.User) *errors.Errors
	GetByEmail(user *domain.User) *errors.Errors
	GetByID(user *domain.User) *errors.Errors
}

type UserRepo struct {
	Db *sql.DB
}

var (
	queryInsertUser     = "INSERT INTO users(first_name,last_name,email,password) VALUES (?,?,?,?);"
	queryGetUserByEmail = "SELECT id, first_name, last_name, password FROM users WHERE email = ? "
	queryGetUserByID    = "SELECT id, first_name, last_name, email FROM users WHERE id = ? "
)

func NewUserRepository(db *db.Adapter) *UserRepo {
	return &UserRepo{Db: db.DB}
}

func (r *UserRepo) Save(user *domain.User) *errors.Errors {
	stmt, err := r.Db.Prepare(queryInsertUser)

	if err != nil {
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	insertResult, saveError := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Password)

	if saveError != nil {
		fmt.Println(saveError.Error())
		return errors.NewInternalServerError(saveError.Error())
	}
	userID, err := insertResult.LastInsertId()

	if err != nil {
		return errors.NewInternalServerError("database error")
	}

	user.ID = userID

	return nil
}

func (r *UserRepo) GetByEmail(user *domain.User) *errors.Errors {
	stmt, err := r.Db.Prepare(queryGetUserByEmail)
	if err != nil {
		return errors.NewInternalServerError("invalid email")
	}

	defer stmt.Close()

	result := stmt.QueryRow(user.Email)

	err = result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Password)
	if err != nil {
		return errors.NewInternalServerError("Failed to login user")
	}

	return nil
}

func (r *UserRepo) GetByID(user *domain.User) *errors.Errors {
	stmt, err := r.Db.Prepare(queryGetUserByID)

	if err != nil {
		return errors.NewInternalServerError("User not found")
	}

	defer stmt.Close()

	result := stmt.QueryRow(user.ID)

	getErr := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)

	if getErr != nil {
		return errors.NewInternalServerError("Couldn't retrieve user")
	}

	return nil
}
