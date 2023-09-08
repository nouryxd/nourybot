package main

import (
	"crypto/tls"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

// Thanks to Twitch moving whispers again I just use email now.
func (app *application) SendEmail() {
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
	m.SetHeader("Subject", "Test Email!")
	m.SetBody("text/plain", "Hello!")

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
