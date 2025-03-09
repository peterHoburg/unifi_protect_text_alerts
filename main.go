package main

import (
	"fmt"
	"net/http"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

func smsNotification(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("sending sms notifications")
	client := twilio.NewRestClient()

	params := &openapi.CreateMessageParams{}
	params.SetFrom(FromPhoneNumber)
	for _, toPhoneNumber := range ToPhoneNumbers {

		params.SetTo(toPhoneNumber)
		params.SetBody("Unifi Protect Alert:")

		_, err := client.Api.CreateMessage(params)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("SMS sent successfully!")
		}
	}
}

func main() {
	initViper()

	http.HandleFunc("/sms-notification", smsNotification)

	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		panic(err)
	}
}
