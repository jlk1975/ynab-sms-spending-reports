/*
package main

import twilio "github.com/twilio/twilio-go"
import openapi "github.com/twilio/twilio-go/rest/api/v2010"
import "os"
import "fmt"

func main() {
    client := twilio.NewRestClient()

    params := &openapi.CreateMessageParams{}
    params.SetTo(os.Getenv("TO_PHONE_NUMBER"))
    params.SetFrom(os.Getenv("TWILIO_PHONE_NUMBER"))
    params.SetBody("Hello from Golang!  💩 😃")

    _, err := client.Api.CreateMessage(params)
    if err != nil {
        fmt.Println(err.Error())
    } else {
        fmt.Println("SMS sent successfully!")
    }
}
*/