# Go Twilio Usage Records

This is a small repository showing the essentials of retrieving Twilio usage records using Go. 

It's not meant to exemplify best practice code, nor idiomatic go.
Rather, it's meant to show the reader, as quickly as possible, how to retrieve Twilio usage data from their account.
Please keep that in mind. 

## Usage

To use the project, first clone the repository, then create a file named _.env_ in the top-level of the newly cloned directory.

In _.env_ add the following code:

```ini
TWILIO_ACCOUNT_SID=<TWILIO_ACCOUNT_SID>
TWILIO_AUTH_TOKEN=<TWILIO_AUTH_TOKEN>
```

After that, from [the Twilio Console's Dashboard](https://twilio.com/console), copy your **Account SID** and **Auth Token** and, in _.env_, replace the respective placeholder values with them (`<TWILIO_ACCOUNT_SID>` and `<TWILIO_AUTH_TOKEN>`).

With _.env_ updated, run the code using the following command.

```bash
go run twilio-usage-records.go
```

You should see output similar to the following.

```bash
+------------------+-----------------+--------------------------+--------+
| START DATE       | END DATE        | CATEGORY                 | PRICE  |
+------------------+-----------------+--------------------------+--------+
| January 28, 2021 | August 26, 2022 | calleridlookups          | $0.00  |
| January 28, 2021 | August 26, 2022 | calls                    | $1.08  |
| January 28, 2021 | August 26, 2022 | calls-client             | $0.00  |
| January 28, 2021 | August 26, 2022 | calls-sip                | $0.00  |
| January 28, 2021 | August 26, 2022 | calls-inbound            | $0.09  |
| January 28, 2021 | August 26, 2022 | calls-inbound-local      | $0.09  |
| January 28, 2021 | August 26, 2022 | calls-inbound-mobile     | $0.00  |
| January 28, 2021 | August 26, 2022 | calls-inbound-tollfree   | $0.00  |
| January 28, 2021 | August 26, 2022 | calls-outbound           | $0.99  |
| January 28, 2021 | August 26, 2022 | phonenumbers             | $15.00 |
| January 28, 2021 | August 26, 2022 | phonenumbers-mobile      | $0.00  |
| January 28, 2021 | August 26, 2022 | phonenumbers-local       | $15.00 |
| January 28, 2021 | August 26, 2022 | phonenumbers-tollfree    | $0.00  |
| January 28, 2021 | August 26, 2022 | shortcodes               | $0.00  |
| January 28, 2021 | August 26, 2022 | shortcodes-customerowned | $0.00  |
| January 28, 2021 | August 26, 2022 | shortcodes-random        | $0.00  |
| January 28, 2021 | August 26, 2022 | shortcodes-vanity        | $0.00  |
| January 28, 2021 | August 26, 2022 | sms                      | $18.02 |
| January 28, 2021 | August 26, 2022 | sms-inbound              | $0.17  |
| January 28, 2021 | August 26, 2022 | sms-inbound-longcode     | $0.17  |
+------------------+-----------------+--------------------------+--------+
| TOTAL RECORDS:   | 20              |                          |        |
+------------------+-----------------+--------------------------+--------+
```