package main

import (
	"crypto/tls"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

// SendEmail sends an email to the email address configured in the .env file
// with the supplied body.
func (app *application) SendEmail(subject, body string) {
	err := godotenv.Load()
	if err != nil {
		app.Log.Fatal("Error loading .env")
	}
	hostname := os.Getenv("EMAIL_HOST")
	login := os.Getenv("EMAIL_LOGIN")
	password := os.Getenv("EMAIL_PASS")
	emailFrom := os.Getenv("EMAIL_FROM")
	emailTo := os.Getenv("EMAIL_TO")
	d := gomail.NewDialer(hostname, 587, login, password)

	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	m := gomail.NewMessage()
	m.SetHeader("From", emailFrom)
	m.SetHeader("To", emailTo)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
