package main

import (
    "fmt"
    "io/ioutil"
	"net/http"
	"github.com/go-chi/chi"
)

func main() {
    port := ":3000"
    r := chi.NewRouter()
    // GET request for buyers
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
    // GET request for products
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
    // GET request for transactions
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
