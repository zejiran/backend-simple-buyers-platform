package handlers

import (
    "fmt"
    "io/ioutil"
	"net/http"
)

func UpdateTransactions() {
    url_transactions := "https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/transactions"
    response, err := http.Get(url_transactions)
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
        fmt.Printf("\n\nWriting on file...")
        data, _ := ioutil.ReadAll(response.Body)
        // Format response to CSV
        raw := string(data)
        processed := ""
        comma_count := 0
        for _, chr := range raw {
            actual := string(chr)
            if actual == "(" {
                processed += "\"["
            } else if actual == ")" {
                processed += "]\""
            } else if actual == "\x00" {
                processed += ","
                comma_count += 1
            } else if actual != "\r" && actual != "#" {
                processed += actual
            }
            if comma_count == 6 {
                processed = processed[:len(processed) - 2]
                processed += "\n"
                comma_count = 0
            }
        }
        // Write to file
        err = ioutil.WriteFile("api-rest/responses/transactions.csv", []byte(processed), 0644)
        fmt.Printf("\n\nFile generated.")
    if err != nil {
        panic(err)
        }
    }
}