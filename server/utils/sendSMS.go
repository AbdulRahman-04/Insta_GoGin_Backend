package utils

import (
	"fmt"

	"github.com/AbdulRahman-04/Go_Backend_Practice/config"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

type SMSData struct {
	From string
	To string
	Body string
}

func SendSMS(data SMSData) error {
	// create a client
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: config.AppConfig.Phone.Sid,
		Password: config.AppConfig.Phone.Token,
	})

	// get the data ready 
	_, err := client.Api.CreateMessage(&openapi.CreateMessageParams{
		From: &config.AppConfig.Phone.Phone,
		To: &data.To,
		Body: &data.Body,
	})
	
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("SMS SENT✅")
	return nil
}