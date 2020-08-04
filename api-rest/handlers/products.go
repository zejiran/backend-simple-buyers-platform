package handlers

import (
    "fmt"
    "io/ioutil"
	"net/http"
	"strings"
)

func UpdateProducts() {
    url_products := "https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/products"
    response, err := http.Get(url_products)
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
        fmt.Printf("\n\nWriting on file...")
        data, _ := ioutil.ReadAll(response.Body)
        csv := strings.Replace(string(data), "'", ",", -1)
        // Write to file
        err = ioutil.WriteFile("api-rest/responses/products.csv", []byte(csv), 0644)
        fmt.Printf("\n\nFile generated.")
        if err != nil {
        panic(err)
        }
    }
}