package main

import (
    "fmt"
    "io/ioutil"
	"net/http"
	"github.com/go-chi/chi"
)

func main() {
    fmt.Println("Starting the application...")
    port := ":3000"
    r := chi.NewRouter()
    /* GET request for buyers
     Parameters:
        - date: date to query on unix timestamp format
     Response:
        Format: JSON
        - id: Buyer ID
        - name: Buyer name
        - age: Buyer age
    */
    r.Get("/buyers", func(w http.ResponseWriter, r *http.Request) {
        url_buyers := "https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/buyers"
        response, err := http.Get(url_buyers)
        if err != nil {
            fmt.Printf("The HTTP request failed with error %s\n", err)
        } else {
            data, _ := ioutil.ReadAll(response.Body)
            w.Write([]byte(string(data)))
        }
    })
    /* GET request for products
     Parameters:
        - date: date to query on unix timestamp format
     Response:
        Format: CSV separed by '
        - id: Product ID
        - name: Product name
        - price: Price in dollar cents
    */
    r.Get("/products", func(w http.ResponseWriter, r *http.Request) {
        url_products := "https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/products"
        response, err := http.Get(url_products)
        if err != nil {
            fmt.Printf("The HTTP request failed with error %s\n", err)
        } else {
            data, _ := ioutil.ReadAll(response.Body)
            w.Write([]byte(string(data)))
        }
    })
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
    r.Get("/transactions", func(w http.ResponseWriter, r *http.Request) {
        url_transactions := "https://kqxty15mpg.execute-api.us-east-1.amazonaws.com/transactions"
        response, err := http.Get(url_transactions)
        if err != nil {
            fmt.Printf("The HTTP request failed with error %s\n", err)
        } else {
            data, _ := ioutil.ReadAll(response.Body)
            w.Write([]byte(string(data)))
        }
    })
    fmt.Println("Serving on localhost" + port)
    http.ListenAndServe(port, r)
}
