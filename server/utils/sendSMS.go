package utils

import (
	"fmt"

	"github.com/AbdulRahman-04/Go_Backend_Practice/config"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

type SmsData struct {
	To string
	From string
	Body string
}

func SendSMS(data SmsData) error {
	// create a clcient 
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: config.AppConfig.Phone.Sid,
		Password: config.AppConfig.Phone.Token,
	})

	// get the message body ready 
	_, err := client.Api.CreateMessage(&openapi.CreateMessageParams{
		To: &data.To,
		From: &config.AppConfig.Phone.Phone,
		Body: &data.Body,
	}) 

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("SMS SENT✅")
	return nil
}