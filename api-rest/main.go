package main

import (
    "fmt"
	"net/http"
	"github.com/go-chi/chi"
    "./handlers"
)

func main() {
    fmt.Println("Starting the application...")
    fmt.Println("----------------------")
    port := ":3717"
    r := chi.NewRouter()
    // Update response files, using by default actual date.
    fmt.Println("Writing on buyers file...")
    handlers.UpdateBuyers()
    fmt.Println("\nWriting on products file...")
    handlers.UpdateProducts()
    fmt.Println("\nWriting on transactions file...")
    handlers.UpdateTransactions()
    // Router
    r.Get("/buyers", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Generating file! Check your api-rest/responses folder."))
    })
    r.Get("/products", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Generating file! Check your api-rest/responses folder."))
    })
    r.Get("/transactions", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Generating file! Check your api-rest/responses folder."))
    })
    // Server
    fmt.Println("----------------\nServing on localhost" + port + "...")
	err := http.ListenAndServe(port, r)
	if err != nil {
		fmt.Println("ListenAndServe:", err)
	}
}
