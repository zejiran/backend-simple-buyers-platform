package main

import (
    "fmt"
	"net/http"
	"github.com/go-chi/chi"
	"time"
	"strconv"
    "./handlers"
    "./database"
)

func main() {
    fmt.Println("Starting the application...")
    fmt.Println("----------------------")
    port := ":3717"
    r := chi.NewRouter()
    // Update response files, using by default actual date.
    currentTime := strconv.FormatInt(time.Now().Unix(), 10)
    fmt.Println("Writing on buyers file...")
    handlers.UpdateBuyers(currentTime)
    fmt.Println("\nWriting on products file...")
    handlers.UpdateProducts(currentTime)
    fmt.Println("\nWriting on transactions file...")
    handlers.UpdateTransactions(currentTime)
    // Upload files to MySQL database
    database.Database()
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
