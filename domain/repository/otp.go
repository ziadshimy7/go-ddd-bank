package repo

import (
	"database/sql"
	"fmt"

	domain "github.com/go-ddd-bank/domain/model"
	"github.com/go-ddd-bank/infrastructure/db"
	errors "github.com/go-ddd-bank/utils"
)

type OTPRepository interface {
	UpdateOTP(otp *domain.OTP) *errors.Errors
	VerifyOTP(otp *domain.OTP) *errors.Errors
}

type OTPRepo struct {
	Db *sql.DB
}

var (
	queryUpdateOTP = `UPDATE otp
	SET
	otp_secret = ?,
	otp_auth_url = ?
	WHERE 
	user_id = ?;`
)

func NewOTPRepository(db *db.DbConnection) *OTPRepo {
	return &OTPRepo{Db: db.DB}
}

func (r *OTPRepo) UpdateOTP(otp *domain.OTP) *errors.Errors {
	stmt, err := r.Db.Prepare(queryUpdateOTP)

	if err != nil {
		fmt.Println(err)
		return errors.NewInternalServerError("User not found")
	}

	defer stmt.Close()

	_, updateErr := stmt.Exec(otp.Otp_secret, otp.Otp_auth_url, otp.ID)

	if updateErr != nil {
		fmt.Println(updateErr.Error())
		return errors.NewInternalServerError("Error executing the update: " + updateErr.Error())
	}
	return nil
}

func (r *OTPRepo) VerifyOTP(otp *domain.OTP) *errors.Errors {
	return errors.NewInternalServerError("cosomk")
}
