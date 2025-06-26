package utils

import (
	"fmt"

	"github.com/AbdulRahman-04/Go_Backend_Practice/config"
	"gopkg.in/gomail.v2"
)

type EmailData struct {
	From string
	To string
	Subject string
	Text string
	Html string
}

func SendEmail(data EmailData) error {
	// get user and pass from config 
	user := config.AppConfig.Email.User
	pass := config.AppConfig.Email.Pass

	// create a new message sender 
	s := gomail.NewMessage()

	s.SetHeader("From", data.From)
	s.SetHeader("To", data.To)
	s.SetHeader("Subject", data.Subject)
	s.SetHeader("Text", data.Subject)
	s.SetHeader("Htmk", data.Html)

	// create an smtp dailer transporter
	t := gomail.NewDialer("smtp.gmail.com", 465, user, pass)

	// send mail 
	if err := t.DialAndSend(s); err !+ nil {
		fmt.Println(err)
	}

	fmt.Println("Email Sent✅")
	return nil
}