// Copyright © 2022 Matthew Setter <matthew@matthewsetter.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// Twilio Usage Records is a small application that supports the
// retrieval of usage records from a Twilio account and renders
// them in a variety of ways.
// At the moment, only tabular format is supported.

package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/bojanz/currency"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/viper"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

// TwilioUsageFormatOptions provides formatting options for
// TwilioUsageTableFormatter and other TwilioUsage*Formatter objects.
type TwilioUsageFormatOptions struct {
	OutputDateFormat string
	InputDateFormat  string
}

// TwilioUsageTableFormatter formats a list of TwilioUsage records in tabular
// format.
type TwilioUsageTableFormatter struct {
	options TwilioUsageFormatOptions
}

// FormatRecords takes an array of Twilio usage records and renders them in a table.
// While there are a number of properties available on a ApiV2010UsageRecord object,
// the function only renders the start and end dates, category, and price in the
// functions output. At some point in the future, this could be changed, but for
// this, small, example, I feel that these four columns suffice3.
func (r *TwilioUsageTableFormatter) FormatRecords(usageRecords []openapi.ApiV2010UsageRecord) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Start Date", "End Date", "Category", "Price"})

	for _, record := range usageRecords {
		price := fmt.Sprintf("%.2f", *record.Price)
		amount, err := currency.NewAmount(price, strings.ToUpper(*record.PriceUnit))
		if err != nil {
			log.Println("Could not create currency amount from usage price.")
			continue
		}
		locale := currency.NewLocale("en_US")
		formatter := currency.NewFormatter(locale)

		startDate, err := time.Parse(r.options.InputDateFormat, *record.StartDate)
		if err != nil {
			log.Println("Could not format start date.")
			continue
		}

		endDate, err := time.Parse(r.options.InputDateFormat, *record.EndDate)
		if err != nil {
			log.Println("Could not format end date.")
			continue
		}

		t.AppendRow([]interface{}{
			startDate.Format(r.options.OutputDateFormat),
			endDate.Format(r.options.OutputDateFormat),
			*record.Category,
			formatter.Format(amount),
		})
	}

	t.AppendSeparator()
	t.AppendFooter(
		table.Row{"Total records: ", len(usageRecords)},
	)

	t.Render()
}

func main() {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	// Add this as a value on a struct
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: viper.GetString("TWILIO_ACCOUNT_SID"),
		Password: viper.GetString("TWILIO_AUTH_TOKEN"),
	})

	//maxRecords := flag.Int("n", 20, "Maximum number of records to return")
	//startDate := flag.String("s", "", "The earliest date of a usage record")
	//endDate := flag.String("e", "", "The latest date of a usage record")

	recordLimit := viper.GetInt("RECORD_LIMIT")
	params := &openapi.ListUsageRecordParams{
		Limit: &recordLimit,
	}
	usageRecords, err := client.Api.ListUsageRecord(params)
	if err != nil {
		panic(fmt.Errorf("Could not retrieve usage records: %w", err))
	}

	recordFormatter := TwilioUsageTableFormatter{
		TwilioUsageFormatOptions{
			viper.GetString("OUTPUT_DATE_FORMAT"),
			viper.GetString("INPUT_DATE_FORMAT"),
		},
	}
	recordFormatter.FormatRecords(usageRecords)
}
