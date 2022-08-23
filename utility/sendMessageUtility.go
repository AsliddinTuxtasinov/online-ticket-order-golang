package utility

import (
	"fmt"
	"net/smtp"
	"os"
)

// This function purpose is send message to email which custom users's
//
// this function allowed three values which toEmailAddress - to email address, subject - for message subject,
// body - for message body.
//
// return error if proces failed
func SendMessageToEmail(toEmailAddress string, subject, body string) error {
	msg := fmt.Sprintf("To: %v\r\nSubject: %v \r\n\r\n%v\r\n", toEmailAddress, subject, body)

	if err := sendMailSmtp([]string{toEmailAddress}, []byte(msg)); err != nil {
		return err
	}

	return nil
}

func sendMailSmtp(to []string, msg []byte) error {

	auth := authSmtp()
	return smtp.SendMail(os.Getenv("SMTP_HOST")+":"+os.Getenv("SMTP_PORT"), auth, os.Getenv("SMTP_FROM"), to, msg)
}

func authSmtp() smtp.Auth {
	return smtp.PlainAuth("", os.Getenv("SMTP_FROM"), os.Getenv("SMTP_PASSWORD"), os.Getenv("SMTP_HOST"))
}
