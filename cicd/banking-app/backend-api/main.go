package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
)

var db = InMemoryDB{
	accounts:     make(map[string]Account),
	transactions: make(map[string]Transaction),
	accountID:    1000,
	transactionID: 5000,
}

func main() {
	// Get port from flag
	addr := flag.String("addr", ":8080", "HTTP network address")
	flag.Parse()

	// Create a new ServeMux
	mux := http.NewServeMux()

	// Register routes
	mux.HandleFunc("GET /accounts", getAccounts)
	mux.HandleFunc("POST /accounts/", createAccount)
	mux.HandleFunc("GET /accounts/{id}", getAccountById)
	mux.HandleFunc("PUT /accounts/{id}", updateAccountById)
	mux.HandleFunc("DELETE /accounts/{id}", deleteAccountById)
	mux.HandleFunc("GET /transactions", getAllTransactions)
	mux.HandleFunc("POST /transactions/", createTransaction)
	mux.HandleFunc("GET /transactions/{id}", getTransactionById)
	mux.HandleFunc("PUT /transactions/{id}", updateTransactionById)
	
	// Health check endpoint
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"status": "healthy"})
	})
	
	// Start server
	fmt.Printf("Banking API Server starting on port %s...\n", *addr)
	log.Fatal(http.ListenAndServe(*addr, mux))
}

