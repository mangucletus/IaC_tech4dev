package main

import "sync"

// InMemoryDB is a simple in-memory database
type InMemoryDB struct {
	accounts     map[string]Account
	transactions map[string]Transaction
	accountID    int
	transactionID int
	mu           sync.RWMutex
}

// Account represents a bank account
type Account struct {
	ID      string  `json:"id"`
	Owner   string  `json:"owner"`
	Balance float64 `json:"balance"`
}

// Transaction represents a banking transaction
type Transaction struct {
	ID          string  `json:"id"`
	AccountID   string  `json:"account_id"`
	Amount      float64 `json:"amount"`
	Type        string  `json:"type"` // "deposit", "withdrawal", "transfer"
	Description string  `json:"description"`
	Status      string  `json:"status"` // "pending", "completed", "failed"
}
