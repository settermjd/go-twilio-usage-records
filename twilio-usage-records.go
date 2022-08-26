package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

func main() {
	params := &openapi.ListUsageRecordParams{}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: os.Getenv("TWILIO_ACCOUNT_SID"),
		Password: os.Getenv("TWILIO_AUTH_TOKEN"),
	})
	usageRecords, err := client.Api.ListUsageRecord(params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		for _, record := range usageRecords {
			fmt.Printf(
				"%s,%s,%s,%f,%s\n",
				*record.StartDate,
				*record.EndDate,
				*record.Category,
				*record.Price,
				*record.PriceUnit,
			)
		}
	}
}
