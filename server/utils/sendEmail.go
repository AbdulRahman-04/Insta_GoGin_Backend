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
	user := config.AppConfig.EMAIL.User
	pass := config.AppConfig.EMAIL.Pass

	// create a seder 
	s := gomail.NewMessage()

	s.SetAddressHeader("From", user, "Team Xpertz")
	s.SetHeader("To", data.To)
	s.SetHeader("Subject", data.Subject)
	s.SetBody("Text/plain", data.Text)
	s.AddAlternative("Html/plain", data.Html)

	// create a transporter
	t := gomail.NewDialer("smtp.gmail.com", 465, user, pass)

	// try send mail
	if err := t.DialAndSend(s); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Email sent✅")
	return nil
}