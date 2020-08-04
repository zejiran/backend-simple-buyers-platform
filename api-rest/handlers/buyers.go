package handlers

import (
    "fmt"
    "io/ioutil"
	"net/http"
)

func UpdateBuyers() {
    url_buyers := "https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/buyers"
    response, err := http.Get(url_buyers)
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
        fmt.Printf("\n\nWriting on file...")
        data, _ := ioutil.ReadAll(response.Body)
        // Write to file
        processed := "{\"buyers\":" + string(data) + "}"
        err = ioutil.WriteFile("api-rest/responses/buyers.json", []byte(processed), 0644)
        fmt.Printf("\n\nFile generated.")
        if err != nil {
            panic(err)
        }
    }
}
