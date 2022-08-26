package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/bojanz/currency"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/joho/godotenv"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

func main() {
	recordLimit := 20
	params := &openapi.ListUsageRecordParams{Limit: &recordLimit}

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
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"Start Date", "End Date", "Category", "Price"})

		for _, record := range usageRecords {
			price := fmt.Sprintf("%.2f", *record.Price)
			amount, err := currency.NewAmount(price, strings.ToUpper(*record.PriceUnit))
			if err != nil {
				fmt.Println("Could not create currency amount from usage price.")
			}
			locale := currency.NewLocale("en_US")
			formatter := currency.NewFormatter(locale)

			startDate, err := time.Parse("2006-01-02", *record.StartDate)
			if err != nil {
				fmt.Println("Could not format start date.")
			}

			endDate, err := time.Parse("2006-01-02", *record.EndDate)
			if err != nil {
				fmt.Println("Could not format end date.")
			}

			t.AppendRow([]interface{}{
				startDate.Format("January 2, 2006"),
				endDate.Format("January 2, 2006"),
				*record.Category,
				formatter.Format(amount),
			})
		}

		t.AppendSeparator()
		t.AppendFooter(table.Row{"Total records: ", len(usageRecords)}, table.RowConfig{AutoMerge: true})

		t.Render()
	}
}
