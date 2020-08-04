package handlers

import (
    "fmt"
    "io/ioutil"
	"net/http"
)

func UpdateBuyers() {
    /* GET request for buyers
     Parameters:
        - date: date to query on unix timestamp format
     Response:
        Format: JSON
        - id: Buyer ID
        - name: Buyer name
        - age: Buyer age
    */
    url_buyers := "https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/buyers"
    response, err := http.Get(url_buyers)
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
        data, _ := ioutil.ReadAll(response.Body)
        // Write to file
        processed := "{\"buyers\":" + string(data) + "}"
        err = ioutil.WriteFile("api-rest/responses/buyers.json", []byte(processed), 0644)
        fmt.Printf("File generated.\n")
        if err != nil {
            panic(err)
        }
    }
}
