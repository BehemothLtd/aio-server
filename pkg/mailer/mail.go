package mailer

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/skratchdot/open-golang/open"
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

	if os.Getenv("MAILER_ADDRESS") != "" {
		if err := Dialer.DialAndSend(m); err != nil {
			return err
		}
	} else {
		// Save the email to a temporary directory
		tempDir := filepath.Join(os.TempDir(), "letter_opener")
		os.MkdirAll(tempDir, os.ModePerm)

		t := time.Now()
		emailFile := filepath.Join(tempDir, fmt.Sprintf("test_email_%s.html", t.Format("20060102150405")))

		f, err := os.Create(emailFile)
		if err != nil {
			fmt.Println("Error creating email file:", err)
			return err
		}
		defer f.Close()

		_, err = f.WriteString(emailMessage.Body)
		if err != nil {
			fmt.Println("Error writing email to file:", err)
			return err
		}

		// Open the email in the default browser
		err = open.Run(emailFile)
		if err != nil {
			return err
		}
	}

	return nil
}
