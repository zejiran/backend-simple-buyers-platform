package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func UpdateProducts(currentTime string) {
	/* GET request for products
	   Parameters:
	      - date: date to query on unix timestamp format
	   Response:
	      Format: CSV separed by '
	      - id: Product ID
	      - name: Product name
	      - price: Price in dollar cents
	*/
	urlProducts := "https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/products"
	response, err := http.Get(urlProducts + "?date=" + currentTime)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		csv := strings.Replace(string(data), "'", ",", -1)
		// Write to file
		err = ioutil.WriteFile("responses/products.csv", []byte(csv), 0644)
		fmt.Printf("File generated.\n")
		if err != nil {
			panic(err)
		}
	}
}
