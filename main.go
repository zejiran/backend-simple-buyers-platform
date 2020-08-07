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
    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        response := "Welcome to root, please go to /buyers, /products or /transactions"
        w.Write([]byte(response))
    })
    r.Get("/buyers", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.Header().Set("Access-Control-Allow-Origin", "*")
        response := database.QueryBuyers()
        w.Write([]byte(response))
    })
    r.Get("/products", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.Header().Set("Access-Control-Allow-Origin", "*")
        response := database.QueryProducts()
        w.Write([]byte(response))
    })
    r.Get("/transactions", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.Header().Set("Access-Control-Allow-Origin", "*")
        response := database.QueryTransactions()
        w.Write([]byte(response))
    })
    // User search
    r.Get("/search/{id}", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.Header().Set("Access-Control-Allow-Origin", "*")
        response := database.SearchBuyer(chi.URLParam(r, "id"))
        w.Write([]byte(response))
    })
    // No search is done
    r.Get("/search/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Write([]byte("[{\"name\": \"Search someone by ID\", \"age\": 0, \"id\": 0}]"))
    })
    // User transactions
    r.Get("/transaction/data/{id}", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.Header().Set("Access-Control-Allow-Origin", "*")
        response := database.TransactionBuyer(chi.URLParam(r, "id"))
        w.Write([]byte(response))
    })
    // Products on transaction
    r.Get("/transaction/products/{id}", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.Header().Set("Access-Control-Allow-Origin", "*")
        response := database.TransactionProducts(chi.URLParam(r, "id"))
        w.Write([]byte(response))
    })
    // Same IP users
    r.Get("/transaction/sameip/{ip}", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.Header().Set("Access-Control-Allow-Origin", "*")
        response := database.SameIPBuyer(chi.URLParam(r, "ip"))
        w.Write([]byte(response))
    })
    // Server
    fmt.Println("----------------\nServing on localhost" + port + "...")
	err := http.ListenAndServe(port, r)
	if err != nil {
		fmt.Println("ListenAndServe:", err)
	}
}
