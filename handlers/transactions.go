package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func UpdateTransactions(currentTime string) {
	/* GET request for transactions
	   Parameters:
	      - date: date to query on unix timestamp format
	   Response:
	      Format: No standard
	      - id: Transaction ID
	      - buyer id: Buyer ID
	      - ip: Buyer IP
	      - device: Buyer devices (ios, android, windows, linux, mac)
	      - product ids: List of purchased products
	*/
	urlTransactions := "https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/transactions"
	response, err := http.Get(urlTransactions + "?date=" + currentTime)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		// Format response to CSV
		raw := string(data)
		processed := ""
		commaCount := 0
		for _, chr := range raw {
			actual := string(chr)
			if actual == "(" {
				processed += "\"["
			} else if actual == ")" {
				processed += "]\""
			} else if actual == "\x00" {
				processed += ","
				commaCount += 1
			} else if actual != "\r" && actual != "#" {
				processed += actual
			}
			if commaCount == 6 {
				processed = processed[:len(processed)-2]
				processed += "\n"
				commaCount = 0
			}
		}
		// Write to file
		err = ioutil.WriteFile("responses/transactions.csv", []byte(processed), 0644)
		fmt.Printf("File generated.\n")
		if err != nil {
			panic(err)
		}
	}
}
