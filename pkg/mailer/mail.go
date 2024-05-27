package mailer

import (
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

var Dialer *gomail.Dialer

type EmailMessage struct {
	From    string
	To      string
	Subject string
	Body    string
}

func NewEmailMessage(to string, subject string) *EmailMessage {
	emailMessage := EmailMessage{
		To:      to,
		Subject: subject,
		From:    os.Getenv("MAILER_FROM"),
	}

	return &emailMessage
}

func InitDialer() *gomail.Dialer {
	if Dialer != nil {
		return Dialer
	}

	port, _ := strconv.Atoi(os.Getenv("MAILER_PORT"))

	Dialer = gomail.NewDialer(
		os.Getenv("MAILER_ADDRESS"),
		port,
		os.Getenv("MAILER_USER"),
		os.Getenv("MAILER_PASSWORD"),
	)

	return Dialer
}

func Send(emailMessage EmailMessage) error {
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("MAILER_FROM"))
	m.SetHeader("To", emailMessage.To)
	m.SetHeader("Subject", emailMessage.Subject)
	m.SetBody("text/html", emailMessage.Body)

	if err := Dialer.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
