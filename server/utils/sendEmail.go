package utils

import (
	"fmt"

	"github.com/AbdulRahman-04/Go_Backend_Project/config"
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
	// get user nd pass
	user := config.AppConfig.Email.User
	pass := config.AppConfig.Email.Pass

	// create a new sender
	s := gomail.NewMessage()

	s.SetAddressHeader("From", user, "Team Instagram")
	s.SetHeader("To", data.To)
	s.SetHeader("Suubject", data.Subject)
	s.SetBody("Text/plain", data.Text)
	s.AddAlternative("Html/plain", data.Html)

	// create a tranporter for smtp
    t := gomail.NewDialer("smtp.gmail.com", 465, user, pass)

	// try sending the mail 
	if err := t.DialAndSend(s); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Email Sent ✅")
    return nil
}
