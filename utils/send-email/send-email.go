package sendemail

import (
	"fmt"
	"net/smtp"
	"os"

	errors "github.com/go-ddd-bank/utils"
)

func SendEmail(msg string, sentTo []string) *errors.Errors {
	from := os.Getenv("MAIL_RU_USERNAME")
	password := os.Getenv("MAIL_RU_PASSWORD")

	smtpHost := "smtp.mail.ru"
	smtpPort := "587"

	message := []byte(msg)

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, sentTo, message)

	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	fmt.Println("Email Sent Successfully!")

	return nil
}
