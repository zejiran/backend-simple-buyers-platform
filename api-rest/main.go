package main

import (
    "fmt"
    "io/ioutil"
	"net/http"
	"github.com/go-chi/chi"
    "strings"
    "unicode"
)

func main() {
    fmt.Println("Starting the application...")
    port := ":3717"
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
            fmt.Printf("\n\nWriting on file...")
            data, _ := ioutil.ReadAll(response.Body)
            w.Write([]byte("File generated! Check your api-rest/responses folder."))
            // Write to file
            processed := "{\"buyers\":" + string(data) + "}"
            err = ioutil.WriteFile("api-rest/responses/buyers.json", []byte(processed), 0644)
            fmt.Printf("\n\nFile generated.")
            if err != nil {
                panic(err)
            }
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
            fmt.Printf("\n\nWriting on file...")
            data, _ := ioutil.ReadAll(response.Body)
            w.Write([]byte("File generated! Check your api-rest/responses folder."))
            csv := strings.Replace(string(data), "'", ",", -1)
            // Write to file
            err = ioutil.WriteFile("api-rest/responses/products.csv", []byte(csv), 0644)
            fmt.Printf("\n\nFile generated.")
            if err != nil {
                panic(err)
            }
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
            fmt.Printf("\n\nWriting on file...")
            data, _ := ioutil.ReadAll(response.Body)
            w.Write([]byte("File generated! Check your api-rest/responses folder."))
            // Format response to CSV
            raw := string(data)
            processed := ""
            first_point := true
            first_time := true
            is_ip := false
            for _, chr := range raw {
                line := ""
                actual := string(chr)
                if first_time {
                    first_time = false
                } else if actual == "#" {
                    line += "\n"
                } else if actual == "." && first_point {
                    line += ","
                    is_ip = true
                    first_point = false
                } else if actual == "." {
                    line += "."
                } else if is_ip && (actual != "." && !unicode.IsDigit(chr)) {
                    is_ip = false
                    line += "," + actual
                } else if actual == "("{
                    line += ",\"["
                } else if actual == ")" {
                    first_point = true
                    line += "]\""
                } else {
                    line += actual
                }
                processed += line
            }
            // Write to file
            err = ioutil.WriteFile("api-rest/responses/transactions.csv", []byte(processed), 0644)
            fmt.Printf("\n\nFile generated.")
            if err != nil {
                panic(err)
            }
        }
    })
    fmt.Println("\nNow go to:\n" +
    "\n http://localhost" + port + "/buyers" +
    "\n http://localhost" + port + "/products" +
    "\n http://localhost" + port + "/transactions" +
    "\n\nThis generates respective files with our responses." +
    "\n\nAfter done, you can exit from this terminal.")
    http.ListenAndServe(port, r)
}
