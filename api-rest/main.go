package main

import (
    "fmt"
	"net/http"
	"github.com/go-chi/chi"
    "./handlers"
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
        handlers.UpdateBuyers()
        w.Write([]byte("Generating file! Check your api-rest/responses folder."))
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
        handlers.UpdateProducts()
        w.Write([]byte("Generating file! Check your api-rest/responses folder."))
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
        handlers.UpdateTransactions()
        w.Write([]byte("Generating file! Check your api-rest/responses folder."))
    })
    fmt.Println("\nNow go to:\n" +
    "\n http://localhost" + port + "/buyers" +
    "\n http://localhost" + port + "/products" +
    "\n http://localhost" + port + "/transactions" +
    "\n\nThis generates respective files with our responses." +
    "\n\nAfter done, you can exit from this terminal.")
    http.ListenAndServe(port, r)
}
