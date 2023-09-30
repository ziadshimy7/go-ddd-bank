package sendemail

import (
	"fmt"
	"net/smtp"

	errors "github.com/go-ddd-bank/utils"
)

func SendEmail(msg string) *errors.Errors {
	from := "ziadshimy77@mail.ru"
	password := "hggTLp4zZ0Xg9nqadXHv"

	to := []string{
		"ziadshimy7@gmail.com",
	}

	smtpHost := "smtp.mail.ru"
	smtpPort := "587"

	message := []byte(msg)

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)

	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	fmt.Println("Email Sent Successfully!")

	return nil
}
