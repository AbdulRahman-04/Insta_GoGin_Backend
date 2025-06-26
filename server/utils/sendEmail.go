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
	// get user nd pass from config 
	user := config.AppConfig.Email.User
	pass := config.AppConfig.Email.Pass

	// create the sender newmessage
	s := gomail.NewMessage()

	s.SetAddressHeader("From", user, "Team Go Pract")
	s.SetHeader("To", data.To)
	s.SetHeader("Subject", data.Subject)
	s.SetBody("Text/plain", data.Text)
	s.AddAlternative("Html/plain", data.Html)
	

	// get smtp transporter ready
	t := gomail.NewDialer("smtp.gmail.com", 465, user, pass)

	// send mail 
	if err := t.DialAndSend(s); err != nil {
		fmt.Println(err)
	}

	fmt.Println("EMAIL SENT✅")
	return nil
}